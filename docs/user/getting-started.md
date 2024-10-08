# Getting Started

The following is an opinionated way of getting started with Flight Control on a local Kind cluster. Please refer to [Installing the Flight Control Service](installing-service.md) and [Installing the Flight Control CLI](installing-cli.md) for the full documentation including other installation options.

## Deploying a Local Kind Cluster

Install the following prerequisites by following the respective documentation:

* `kind` latest version ([installation guide](https://kind.sigs.k8s.io/docs/user/quick-start/))
* `kubectl` CLI of a matching version ([installation guide](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/))
* `helm` CLI version v3.15 or later ([installation guide](https://helm.sh/docs/intro/install/))

Deploy a Kind cluster:

```console
$ kind create cluster

enabling experimental podman provider
Creating cluster "kind" ...
[...]
```

Verify the cluster is up and you can access it:

```console
$ kubectl get pods -A

NAMESPACE            NAME                                         READY   STATUS    RESTARTS   AGE
kube-system          coredns-76f75df574-v6plv                     1/1     Running   0          49s
kube-system          coredns-76f75df574-xfm2w                     1/1     Running   0          49s
kube-system          etcd-kind-control-plane                      1/1     Running   0          61s
kube-system          kindnet-kznkx                                1/1     Running   0          49s
kube-system          kube-apiserver-kind-control-plane            1/1     Running   0          61s
kube-system          kube-controller-manager-kind-control-plane   1/1     Running   0          61s
kube-system          kube-proxy-qffqj                             1/1     Running   0          49s
kube-system          kube-scheduler-kind-control-plane            1/1     Running   0          65s
local-path-storage   local-path-provisioner-7577fdbbfb-wxbck      1/1     Running   0          49s
```

Verify Helm is installed and can access the cluster:

```console
$ helm list

NAME  NAMESPACE  REVISION  UPDATED  STATUS  CHART  APP VERSION
```

## Deploying the Flight Control Service

### Standalone flightctl with keycloak integration

Create a values.yaml file with the following content, replace flightctl.MY.DOMAIN with your base
domain. Please note this values file will be simplified in the future to avoid duplication
by making use of the global.flightctl.baseDomain value.

```yaml
global:
  flightctl:
    baseDomain: "flightctl.MY.DOMAIN"
    clusterLevelSecretAccess: true
    useRoutes: true
  storageClassName: "lvms-vg1"
flightctl:
  api:
    auth:
      oidcAuthority: "https://auth.flightctl.MY.DOMAIN/realms/flightctl"
      internalOidcAuthority: "http://keycloak:8080/realms/flightctl"
      enabled: true

# using keycloak as our OIDC authority for authentication

keycloak:
  enabled: true
  namespace: "flightctl"
  db:
    namespace: "flightctl"

  realm:
    redirectUris:
      - /realms/flightctl/account/*
      - http://127.0.0.1/callback
      - https://ui.flightctl.MY.DOMAIN/*
      - https://ui.flightctl.MY.DOMAIN/
      - https://ui.flightctl.MY.DOMAIN
    webOrigins:
      - https://api.flightctl.MY.DOMAIN
      - https://ui.flightctl.MY.DOMAIN
    adminUrl: "https://auth.flightctl.MY.DOMAIN"
    baseUrl: "https://auth.flightctl.MY.DOMAIN"
    rootUrl: "https://auth.flightctl.MY.DOMAIN"

  # section consumed by the ui charts
  # using keycloak as our OIDC authority for authentication
  authority: https://auth.flightctl.MY.DOMAIN/realms/flightctl
  clientid: flightctl
  redirect: https://ui.flightctl.MY.DOMAIN

# ui configuration
flightctlUi:
  namespace: flightctl
  hostName: ui.flightctl.MY.DOMAIN
  image: quay.io/flightctl/flightctl-ui:latest
  flightctlServer: https://flightctl-api:3443
  flightctlMetricsServer: https://flightctl-api:9090
  bootcImgUrl: quay.io/example/example-agent-centos:bootstrap
  qcow2ImgUrl: https://example.com/disk.qcow2
  certs:
    ca: "" # use --set-file flightctlUi.certs.ca=ca.crt
    frontRouter: ""
    frontRouterKey: ""

```

Install a released version of the Flight Control Service into the cluster by running:

```console
$ helm upgrade --install --version=0.1.1 \
    --namespace flightctl --create-namespace \
    flightctl oci://quay.io/flightctl/charts/flightctl \
    --values values.yaml

```

Retrieve the CA certificate for the API service:

```console
kubectl rollout status deployment flightctl-api -n flightctl -w --timeout=300s

API_POD=$(kubectl get pod -n flightctl -l flightctl.service=flightctl-api --no-headers -o custom-columns=":metadata.name" | head -1)

kubectl exec -n flightctl "${API_POD}" -- cat /root/.flightctl/certs/ca.crt > ca.crt
```

Install a release version of the Flight Control UI into the cluster by running:

```console
$ helm upgrade --install --version=0.1.0 \
    --namespace flightctl --create-namespace \
    flightctl-ui oci://quay.io/flightctl/charts/flightctl-ui \
    --values values.yaml \
    --set-file flightctlUi.certs.ca=ca.crt
```

Verify your Flight Control Service is up and running:

```console
$ kubectl get pods -n flightctl

[...]
```

## Installing the Flight Control CLI

In a terminal, select the appropriate Flight Control CLI binary for your OS (linux or darwin) and CPU architecture (amd64 or arm64), for example:

```console
$ FC_CLI_BINARY=flightctl-linux-amd64

[...]
```

Download the `flightctl` binary to your machine:

```console
$ curl -LO https://github.com/flightctl/flightctl/releases/download/latest/${FC_CLI_BINARY}

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 29.9M  100 29.9M    0     0  5965k      0  0:00:05  0:00:05 --:--:-- 7341k
```

Verify the downloaded binary has the correct checksum:

```console
$ echo "$(curl -L -s https://github.com/flightctl/flightctl/releases/download/latest/${FC_CLI_BINARY}-sha256.txt)  ${FC_CLI_BINARY}" | shasum --check

flightctl-linux-amd64: OK
```

If the checksum is correct, rename it to `flightctl` and make it executable:

```console
$ mv "${FC_CLI_BINARY}" flightctl && chmod +x flightctl

[...]
```

Finally, move it into a location within your shell's search path.

## Logging into the Flight Control Service from the CLI

Retrieve the password for the "demouser" account that's been automatically generated for you during installation:

```console
$ kubectl get secret/keycloak-demouser-secret -n flightctl -o=jsonpath='{.data.password}' | base64 -d

[...]
```

Use the CLI to log into the Flight Control Service:

```console
$ flightctl login https://api.flightctl.127.0.0.1.nip.io/ --web --insecure-skip-tls-verify

[...]
```

In the web browser that opens, use the login "demouser" and the password you retrieved in the previous step.

Verify you can now access the service via the CLI:

```console
$ flightctl get devices

NAME                                                  OWNER   SYSTEM  UPDATED     APPLICATIONS  LAST SEEN
```

## Login into the Flight Control Service from the standalone UI

Browse to `ui.flightctl.MY.DOMAIN` and login with the demouser obtained from the previous step.

## Building a Bootable Container Image including the Flight Control Agent

Next, we will use [Podman](https://github.com/containers/podman) to build a [bootable container image (bootc)](https://containers.github.io/bootc/) that includes the Flight Control Agent binary and configuration. The configuration contains the connection details and credentials required by the agent to discover the service and send an enrollment request to the service.

Retrieve the agent configuration with enrollment credentials by running:

```console
$ flightctl certificate request --scope=enrollment --validity=1y -o agent-config > config.yaml

[...]
```

The returned `config.yaml` should look similar to this:

```console
$ cat config.yaml

enrollment-service:
  service:
    server: https://agent-api.flightctl.127.0.0.1.nip.io:7443
    certificate-authority-data: LS0tLS1CRUdJTiBD...
  authentication:
    client-certificate-data: LS0tLS1CRUdJTiBD...
    client-key-data: LS0tLS1CRUdJTiBF...
  enrollment-ui-endpoint: https://ui.flightctl.127.0.0.1.nip.io:8080
```

Create a `Containerfile` with the following content:

```console
$ cat Containerfile

FROM quay.io/centos-bootc/centos-bootc:stream9

RUN dnf -y copr enable @redhat-et/flightctl-dev centos-stream-9-x86_64 && \
    dnf -y install flightctl-agent; \
    dnf -y clean all; \
    systemctl enable flightctl-agent.service

ADD config.yaml /etc/flightctl/
```

Note this is a regular `Containerfile` that you're used to from Docker/Podman, with the difference that the base image referenced in the `FROM` directive is bootable. This means you can use standard container build tools and workflows.

For example, as a user of Quay who has the privileges to push images into the `quay.io/${YOUR_QUAY_ORG}/centos-bootc-flightctl` repository, build the bootc image like this:

```console
$ sudo podman build -t quay.io/${YOUR_QUAY_ORG}/centos-bootc-flightctl:v1

[...]
```

Log in to your Quay account:

```console
$ sudo podman login quay.io

Username: ******
Password: ******
Login Succeeded!
```

Push your bootc image to Quay:

```console
$ sudo podman push quay.io/${YOUR_QUAY_ORG}/centos-bootc-flightctl:v1

[...]
```

## Provisioning a Device with a Bootable Container Image

A bootc image is a file system image, i.e. it contains the files to be written into an existing file system, but not the disk layout and the file system itself. To provision a device, you need to generate a full disk image containing the bootc image.

Use the [`bootc-image-builder`](https://github.com/osbuild/bootc-image-builder) tool to generate that disk image as follows:

```console
$ mkdir -p output && \
  sudo podman run --rm -it --privileged --pull=newer --security-opt label=type:unconfined_t \
    -v $(pwd)/output:/output -v /var/lib/containers/storage:/var/lib/containers/storage \
    quay.io/centos-bootc/bootc-image-builder:latest \
    --type raw quay.io/${YOUR_QUAY_ORG}/centos-bootc-flightctl:v1

[...]
```

Once `bootc-image-builder` completes, you'll find the raw disk image under `output/image/disk.raw`. Now you can flash this image to a device using standard tools like [arm-image-installer](https://docs.fedoraproject.org/en-US/iot/physical-device-setup/#_scripted_image_transfer_with_arm_image_installer), [Etcher](https://etcher.balena.io/), or [`dd`](https://docs.fedoraproject.org/en-US/iot/physical-device-setup/#_manual_image_transfer_with_dd).

For other image types like QCoW2 or VMDK or ways to install via USB stick or network, see [Building Images](building-images.md).

## Enrolling a Device

When the Flight Control Agent first starts, it sends an enrollment request to the Flight Control Service. You can see the list of requests pending approval with:

```console
$ flightctl get enrollmentrequests

NAME                                                  APPROVAL  APPROVER  APPROVED LABELS
54shovu028bvj6stkovjcvovjgo0r48618khdd5huhdjfn6raskg  Pending   <none>    <none>    
```

You can approve an enrollment request and optionally add labels to the device:

```console
$ flightctl approve -l region=eu-west-1 -l site=factory-berlin 54shovu028bvj6stkovjcvovjgo0r48618khdd5huhdjfn6raskg

Success.

$ flightctl get enrollmentrequests

NAME                                                  APPROVAL  APPROVER  APPROVED LABELS
54shovu028bvj6stkovjcvovjgo0r48618khdd5huhdjfn6raskg  Approved  demouser  region=eu-west-1,site=factory-berlin
```

After the enrollment completes, you can find the device in the list of devices:

```console
$ flightctl get devices

NAME                                                  OWNER   SYSTEM  UPDATED     APPLICATIONS  LAST SEEN
54shovu028bvj6stkovjcvovjgo0r48618khdd5huhdjfn6raskg  <none>  Online  Up-to-date  <none>        3 seconds ago
```

## Where to go from here

Now that you have a Flight Control-managed device, refer to [Managing Devices](managing-devices.yaml) and [Managing Fleets](managing-fleets.yaml) for how to configure and update individual or fleets of devices.

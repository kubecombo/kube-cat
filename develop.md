# [operator-sdk](https://sdk.operatorframework.io/docs/installation/#install-from-github-release)

```bash

# install operator-sdk

export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac)
export OS=$(uname | awk '{print tolower($0)}')

export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/v1.41.1
curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}

mv operator-sdk_linux_amd64 /usr/bin/operator-sdk
chmod +x /usr/bin/operator-sdk

```

## 1. code init

init project

``` bash

operator-sdk init --plugins go/v4 --domain kubecat.com --owner "kubecat" --repo github.com/kubecombo/kube-cat

```

create api

> CRD 命名尽量至少由两个单词组成

``` bash

operator-sdk create api --group cat --version v1 --kind ClusterChecker --resource --controller
operator-sdk create api --group cat --version v1 --kind OneChecker --resource --controller
operator-sdk create api --group cat --version v1 --kind CustomPing --resource --controller

go mod tidy
make generate
make manifests

```

init webhook

```bash

operator-sdk create webhook --group cat --version v1 --kind ClusterChecker --defaulting --programmatic-validation
operator-sdk create webhook --group cat --version v1 --kind OneChecker --defaulting --programmatic-validation
operator-sdk create webhook --group cat --version v1 --kind CustomPing --defaulting --programmatic-validation

```

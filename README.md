# `stemcell-diff`

Install:

```bash
go get github.com/govau/stemcell-diff
```

Usage:

```bash
stemcell-diff xenial-97.18 xenial-97.28
```

Output:

```text
                                       xenial-97.18          xenial-97.28
apparmor                               2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
apparmor-utils                         2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
initramfs-tools                        0.122ubuntu8.12       0.122ubuntu8.13
initramfs-tools-bin                    0.122ubuntu8.12       0.122ubuntu8.13
initramfs-tools-core                   0.122ubuntu8.12       0.122ubuntu8.13
libapparmor-perl                       2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
libapparmor1:amd64                     2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
linux-generic-hwe-16.04                                      4.15.0.38.61
linux-generic-hwe-16.04-edge           4.15.0.34.55
linux-headers-4.15.0-34                4.15.0-34.37~16.04.1
linux-headers-4.15.0-34-generic        4.15.0-34.37~16.04.1
linux-headers-4.15.0-38                                      4.15.0-38.41~16.04.1
linux-headers-4.15.0-38-generic                              4.15.0-38.41~16.04.1
linux-headers-generic-hwe-16.04                              4.15.0.38.61
linux-headers-generic-hwe-16.04-edge   4.15.0.34.55
linux-image-4.15.0-34-generic          4.15.0-34.37~16.04.1
linux-image-4.15.0-38-generic                                4.15.0-38.41~16.04.1
linux-image-generic-hwe-16.04                                4.15.0.38.61
linux-image-generic-hwe-16.04-edge     4.15.0.34.55
linux-libc-dev:amd64                   4.4.0-135.161         4.4.0-138.164
linux-modules-4.15.0-34-generic        4.15.0-34.37~16.04.1
linux-modules-4.15.0-38-generic                              4.15.0-38.41~16.04.1
linux-modules-extra-4.15.0-34-generic  4.15.0-34.37~16.04.1
linux-modules-extra-4.15.0-38-generic                        4.15.0-38.41~16.04.1
openssh-client                         1:7.2p2-4ubuntu2.4    1:7.2p2-4ubuntu2.5
openssh-server                         1:7.2p2-4ubuntu2.4    1:7.2p2-4ubuntu2.5
openssh-sftp-server                    1:7.2p2-4ubuntu2.4    1:7.2p2-4ubuntu2.5
python3-apparmor                       2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
python3-libapparmor                    2.10.95-0ubuntu2.9    2.10.95-0ubuntu2.10
tzdata                                 2017c-0ubuntu0.16.04  2018e-0ubuntu0.16.04
```

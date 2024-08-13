det_key
=======

Generate deterministic ed25519 SSH keys from a seed file.

Usage:
```
$ go build .
$ echo "HELLO" > seed
$ ./det_key seed foo foo.pub
PEM private key written to: foo
Public key written to: foo.pub
$ cat foo
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtz
c2gtZWQyNTUxOQAAACCgn5LATtPR06Wxc4OL1A/1sci7Yxi86nD09yqL/w7CoQAA
AIjLSSnSy0kp0gAAAAtzc2gtZWQyNTUxOQAAACCgn5LATtPR06Wxc4OL1A/1sci7
Yxi86nD09yqL/w7CoQAAAEDYxbXhMNFpTmWxvz+RUCTVHoeBckirYl6HMuGDwyGp
aqCfksBO09HTpbFzg4vUD/WxyLtjGLzqcPT3Kov/DsKhAAAAAAECAwQF
-----END OPENSSH PRIVATE KEY-----
$ ssh-keygen -y -f ./foo
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKCfksBO09HTpbFzg4vUD/WxyLtjGLzqcPT3Kov/DsKh
$ cat foo.pub
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKCfksBO09HTpbFzg4vUD/WxyLtjGLzqcPT3Kov/DsKh
$ rm foo foo.pub
$ ./det_key seed foo foo.pub
PEM private key written to: foo
Public key written to: foo.pub
$ cat foo.pub
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKCfksBO09HTpbFzg4vUD/WxyLtjGLzqcPT3Kov/DsKh
$ cat foo
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtz
c2gtZWQyNTUxOQAAACCgn5LATtPR06Wxc4OL1A/1sci7Yxi86nD09yqL/w7CoQAA
AIiZdOFXmXThVwAAAAtzc2gtZWQyNTUxOQAAACCgn5LATtPR06Wxc4OL1A/1sci7
Yxi86nD09yqL/w7CoQAAAEDYxbXhMNFpTmWxvz+RUCTVHoeBckirYl6HMuGDwyGp
aqCfksBO09HTpbFzg4vUD/WxyLtjGLzqcPT3Kov/DsKhAAAAAAECAwQF
-----END OPENSSH PRIVATE KEY-----
```

This tool is for illustration purposes only and has not been peer
reviewed by anyone. Do not use in production.
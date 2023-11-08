# A Porter Spin Mixin
This is a Spin mixin for Porter. It (at this time) focuses on deploying Spin applications to Fermyon or Fermyon Cloud. 
If Fermyon Cloud, set the Fermyon Cloud boolean to true:
```
mixins:
  - spin:
    fermyonCloud: true
```
If you have a self-installed Fermyon, then you will need the following information to connect: URL, Hippo Admin, Hippo Password, and Bindle URL.

By default Porter copies all the files in the bundle directory into the invocation image. If your necessary Spin files are not in the bundle directory, then
you can specify that with `workingDir` tag. This can be used within the action area of the bundle, or underneath the mixin:
```
mixins:
  - spin:
    workingDir: "~/mysecret/spinapp"
```

```
install:
   spin:
    workingDir: "~/mysecret/spinapp"
```


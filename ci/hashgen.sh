
#!/bin/sh

for f in kube-sync*; 
    do shasum -a 256 "$f" > "$f".sha256sum; 
done

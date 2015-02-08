#!/bin/bash

## double colon bash only, cannot use for sh
## eg.  
##    OK -> $./double-colon.sh   $ bash ./double-colon.sh
##    NG -> $ sh double-colon.sh 
##          double-colon.sh: line 11: `hello::foo': not a valid identifier

hello::foo() {
  (
    echo "hello,foo"
  )	
}

hello::foo

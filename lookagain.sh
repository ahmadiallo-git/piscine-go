#!usr/bin/bash
find -iname '*.sh' -printf '%f\n' | cut -d '.' -f1 | cut -d '\' -f2
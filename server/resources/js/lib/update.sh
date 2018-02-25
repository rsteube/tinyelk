#!/bin/bash
set -e

declare -A pkg
pkg[c3]='0.4.21'
pkg[d3]='3.5.17'
pkg[mithril]='1.1.6'

rm -f *.js
for i in "${!pkg[@]}"; do
  package=$i
  version=${pkg[$i]}

  curl -s "https://cdnjs.cloudflare.com/ajax/libs/${package}/${version}/${package}.min.js" > "${package}.js"
done


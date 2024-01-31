#!/bin/bash
for i in {0..1}
do 
  cp bluechi-template.yaml "${i}-bluechi.yaml"
  sed -i "s/{{id}}/$i/g" "${i}-bluechi.yaml"
done

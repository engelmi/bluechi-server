#!/bin/bash
for i in {0..10000}
do 
  file="${i}-bluechi.yaml"
  kubectl apply -f "${file}"
done

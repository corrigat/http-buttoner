#!/bin/bash

thistr="Blocking on: $(date)"
echo "${thistr}" >> dnstatus.txt
echo "{\"status\": \"${thistr}\"}" > static/current.txt


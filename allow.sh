#!/bin/bash

thistr="Allowing on: $(date)"
echo "${thistr}" >> dnstatus.txt
echo "{\"status\": \"${thistr}\"}" > static/current.txt
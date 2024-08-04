#!/bin/bash

PROJECT=homeopathic-doctor-assistant
local_import=github.com/hamza-sharif/${PROJECT}

go build -buildvcs=false -o "bin/${PROJECT}-server" "${local_import}/cmd"

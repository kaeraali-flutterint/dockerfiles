# Dockerfiles

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/kaeraali-flutterint/dockerfiles/ci.yml)

> Another sandbox repo for testing how to build docker images

Features:

- Only builds the images you've actually changed
- Runs terratest tests for docker images on PRs
- Publishes images on merge to main

Missing features:

- Ability to authenticate to anything other than ghcr.io
- Authenticated pull of base images
- Security scanning
- Ability to publish tagged images from a branch
- Re-run builds on changes to github workflow

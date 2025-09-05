# CRE - Common Reliability Enumerations

An open standard for naming, categorizing, and detecting reliability problems

[Documentation](https://docs.prequel.dev/cres/introduction) | [Slack](https://inviter.co/prequel) | [Playground](https://play.prequel.dev/) | [Mailing List](https://www.detect.sh)

[![Unit Tests](https://github.com/prequel-dev/cre/actions/workflows/build.yml/badge.svg)](https://github.com/prequel-dev/cre/actions/workflows/build.yml)
[![Unit Tests](https://github.com/prequel-dev/preq/actions/workflows/build.yml/badge.svg)](https://github.com/prequel-dev/preq/actions/workflows/build.yml)
[![Unit Tests](https://github.com/prequel-dev/prequel-compiler/actions/workflows/build.yml/badge.svg)](https://github.com/prequel-dev/prequel-compiler/actions/workflows/build.yml)

---

## Overview

### What are CREs?

***Common Reliability Enumerations (CREs)*** are an open, structured standard for naming and categorizing reliability problems found in production systems. CREs represent the collective knowledge of [The Open Problem Detection (and Resolution) Community](https://detect.sh) where hundreds of engineers and practitioners across startups, enterprises, and critical infrastructure providers discuss how to share, detect, and mitigate reliability problems.

CREs provide a consistent way to describe reliability problems (cause, impact, and mitigation). The CRE schema and taxonomy enables the sharing of reliability intelligence and gives teams a vocabulary to discuss recurring problems without reinventing the wheel or diagnosing incidents in isolation.

Just as CVEs (Common Vulnerabilities and Exposures) provide a method to classify and share known threats, CREs offer an equivalent standard for reliability problems.

With CREs, you can:

* Recognize known failure modes before they escalate
* Correlate similar issues across services, teams, or companies
* Drive better postmortems, triage, and tooling decisions
* Contribute your own findings to an evolving, community-backed index

CREs give teams a common framework to ***identify***, ***compare***, and ***learn*** from reliability issues—making patterns visible that were previously siloed or overlooked.

## Getting Started

### Schema

The Common Relability Enumeration Schema is located in [cre-schema.json](cre-schema.json). Learn more about the [CRE specification](http://docs.prequel.dev/cres/cre-spec) and [rule syntax](http://docs.prequel.dev/rules/syntax).

### CRE Rules

* CRE rules are located in the [`rules/`](rules/) folder. Each CRE is placed in its own folder. 
* Tags and categories are also located in this folder in the [`rules/tags`](rules/tags/) subfolder.

### Rule Builder

A CRE builder tool `ruler` [is provided](https://github.com/prequel-dev/cre/releases) to validate CREs and generate a final rules document for a problem detector to consume. The rule builder generates and adds rule hashes derived from the content of the rules. The rule hash will only change if the content of the rule changes. It also validates tag and category references and ensures there are no duplicate IDs.

Check out [CONTRIBUTING.md](CONTRIBUTING.md) to learn how to build and test your first rule.

### Playground

The fastest way to quickly test a rule on data is with the [CRE playground](https://play.prequel.dev/?rule=E4VwNgpgzgXAUAAgQWgQY2BeScIJYAmMCAwgEoCiyATAAzUAsytLA7IrglBAG4TB4ALgE9itDrkFDIxMgEMARgqEBZAIoIVAO2h45CAPZ9gYA3IIQCCTGiP88WgOYIADvyh4ogiFsEIAjiAQQVASOGhy3o4GwKIIALbQUHKOEMiBwWkuwAYKkPGhnAhyIIIAFjHEAAqYGWBhSBZQGHguUgZaxAA%2BDTioACplEAjySqoaaGAgXvz4UK45aEkOzvpgcsCpCFog8QqzBgBmru6e3r4JeMA5mFYZIcV%2BCgYGggB0CL0R01gIPUVIAZDBAUExyJwLAxLKBQAA0mh0Hjk8M8hmMpnMlgQAHcynhINYILZjCsTsAPDMLvdoIYtAhnq8Pr08PEXHI0IJur1ASNFMpBOo5ggQFpFATBAZIdD5uDhNsINiEkkUjTwVYInTvtxDKUVfMHOgOlBdu5ilorNkDAQQEtyW9eoIUrBPgCUOhMDR6EwWLR2K7UABrLQGbFaZCWvIQeLct3APlCeL%2BXrxIR4RyRPAdLn%2BhAAQQIACtplIISnrjFSS4DGA8Gg8DSJQgaym-OVhjs9gdjmWbljqYUAagyFG7Ag8Y4ysg5Dw5PjFPihHKqzW6zTDjl4gEgiEY6h81ZzARUx05GBSFUAKqEqAGEDAaFmqyYLwbVvA0b8wWTYv8XcIC-agA2usABecr9gAugAFGUgiCC4sAAPSIdiqFvHGYyCImby2PEiEEFCUCIaBwjpNu0AAJQII204GIQ%2BBaGgd4CBCbaGl48xHDiAgls4BCRPojZHlAAYupwmCHPwPjQtgg5jnBCEwMhjg5CACFvNEBiOJAOEGHhjiIRh-KJsgPzkohaCIRAAYAGoAJyCP0ABCZAAFoFIheEAIwUAAGtiyCgCATlqLmABSvRyC4LgrhmRqyUUqCiokxAAERGQm-ipTGSDGB4WYIKlADMbx2W8AAe2W4Ikjr8Y6CU4AGDhEAsEB1MyLUAKwXgAIl5bnlQA4iQnUADJhWoWg2S4ua2V5Xm5r0qQ6HG7SdAgXlhKAMi9NwGSMVgMbYs1IbEEVtADkUvA%2BJyOVcLe96-BlWFJq6MQWMADWJYSqTlWlPWeBEwBHqWyqpFBbwANQUQa4KGGAVgOEDoprRD0NceUqLBhYVWum6qXaLo%2BionYJhmBYBC40UOjpt4X2cKgqUAMoAJKDf0FBkCohJLHgfBWKgUBlKUvEIARoapUAA&data=EwBmFYFoQZkhGeAVeAWAXCE74E4B0WRxAWgAQDOALgCYD2ArlWQGJmgTRyJlqbZ58qYPAAcANnABqLPzIBtAKYAnZXQC6ZADwh8wYARAA%2BMgBEAlhQDGAQ2U1zAOwDmZALaKKFG88VkA3gDkACS%2BjgD6ttSBADT%2BAGZ0ylaK4YoAboqOVOHKivF5FAAWMQDEAEr5OkKouOAwIOKiqPiIAOy4uPAgaG1C4m31RgC%2Bw2QFdG7auvqGJlR00-gw8OBzZE5kNo5kdAA2NBuOtsqONlTmdDsAFPBtqPDiIOCoMKgAlLvxZFRFlmSOOg0Py3e53HriVafLQAQSMACgOFBYAhkHxZKBCMRSJRaIxmGwkVxUbwMLJBMIxG1gDJsFgFCo1JpqrNCCYLNY7A4XO5PN5fAEQmFIjZonFEslUhksjk8gVPCUKlVdKhavVGs1WndOt1ev1BjARmMJlMWQY2T9FtUVuBxBbNttdgcjiczhcrmRQQ8ni83p86N9fv9AcDPXcHm0IVDEWBkdw0WTsCIsdisORqPQmKx2LHiTx0QICG0xBJRLS5PInIlmbpwIJjGQAMpUOxURSHDxeHx%2BahJPwBn4ATwADn5R8oKJY29lxkkyOkinRqGRAgB6QIxzgoxAoRPofQp1PpvFZwm57fwUn8fe6NBq8Tl%2BnyADudkcNb0qD6DYAso5POYNjXIEyg2AARmB5hUAAAqBEFQW4ACO0D4BQ6RWPgVh7Aw1AqPgex0LYeyBO86BkAAVORZAAOowuUAByACS9EAOIUVRf4ATYGwULsmTKARNjAjQZH%2BDQDBuMO4QEc4MTPsoUGpL8hSLgcwybnGqK7tewAwIe2LHpmBI5lu8ZXhit5tMWuCoI%2BIAKK%2BpwfsAX4Wpxk5ASB4GQTBcE%2BUhKFoRhWE4W2yj4YRNjEaR7E0XRTGsbF7mATxfEqIJwmieJknSXQsnyYp4TKQq%2Bw0OpRIXtpGItKmRCGfi2YVWZBY3nodyiIgdkOW%2Bzmub%2B-4ecBflQbB3kIchuhBZh2G4eFBFESRZGUXFDHMWxy3Jdx-x0PxGXtllElSTJckKW2RVFCppXlee8ZVUm4D6TiGYNWepkki1mIuequBdS%2BPVLC534mJtQ1jb5YMBZN6HTaFeHzVFi2xbRq2JRtA0pdtu10EJ%2B0BNlR15SdhXFcUV0aXmCY6XatVpriRmNTd717p94i4DAwAPrIT6Oe%2BAN9cD6OecN4PwVQkOodDIWzRFC0xctyMJetHGC6lO3pdjmV44duX5adSkXSVank5VH3fjT9WniZmn5szMx1HcbS-TzvVA2QINeaLo2i%2BLU1S2FMsI3LVEK2tSUq5j6s4yJWs5cdBVnSTqllcbt0faIj11XTL1WxT5lJroKyiKI9xO-9LL827gug57ws%2B5LM3%2B-D0VLcH8Wh2jXGq1jUcHbHhPx-rl1G01WkfYY5tZ5bI82zpBcPAY8Cl05fOu%2B7wte-5E0S8FDdw5FzdI23qPK53EcCRruNidrcd6%2BdQ-J9PlOyA0Ge089U%2BMzPGLwMsMBtKIIhfpVg0EsOs6xKhWDVgpHkqAXj2UQgwRQiDeIDioCOPwwtIh7FFJOKw4QEFIL8AsOgABrXgog6xuAoCnUee4VivxABbYyj8877h-jAP%2BNkQBAMcNWUB9YTAQKgU4Vw8DEHIK%2BEOUcZBMEIKSBJfB4iiF0FIWQEAVCaE7gLMsYADCmEMzel-JM7CYCiBAKgTmdJ7KVl4SA6oYCLRCP4iItRZACESNQegmRYNwjUDyDYNwijCGWlUeo6hHBxAUzuugDhejJ7MM-peD6P9wDPEkLZLm1jgHOXNA2coDBHCOBcWBFRzBcLDjIIoSSaDwg0DApEIoigrBkOBPEJw7YyBgUHFsYcFThZA). The playground runs as [WebAssembly (wasm)](https://web.dev/explore/webassembly) in the browser. Data and rules are not sent to an API. No data leaves your browser.

### Problem Detector

[preq](https://github.com/prequel-dev/preq) is a free and open community-driven reliability problem detector that runs CREs on data. Use it to develop and test CREs on Linux, macOS, or Windows.

## How to contribute

New contributors are encouraged to join the problem detection community add new CREs. Learn how to contribute in [CONTRIBUTING.md](CONTRIBUTING.md).

## Rule Coverage

### Tags & Categories

CREs are organized by tags and categories.

* [Tags](https://docs.prequel.dev/cres/public?view=tags)
* [Categories](https://docs.prequel.dev/cres/public?view=categories)

### Technology Coverage

CREs exist for both popular and obscure project. 

* [CREs by Technology](https://docs.prequel.dev/cres/public?view=technologies)


## Join the community!

* [Discussions](https://github.com/prequel-dev/cre/discussions)
* [Slack](https://prequel-dev.slack.com/)
* [Mailing list](https://www.detect.sh/)

# Contributing to CREs

Thank you for your interest in contributing to Common Reliability Enumerations (CREs)! We appreciate your efforts to improve the problem detection community. Before you get started, please review these guidelines to help ensure your contributions are effective and easy to understand.

## Why We Create Issues Before Contributing Code or New Rules

We encourage everyone to open an issue on GitHub before submitting any code or new rules. Doing so:

* Helps us discuss and refine ideas early.
* Ensures the proposed change works well in various environments.
* Saves time by reducing the risk of major rework later on.

Starting with a pull request before discussing it can be disappointing if we realize that the rule or approach needs significant changes. Instead, opening an issue first lets us brainstorm together and make sure we’re on the right track from the start!

### What a good issue looks like

We offer a [few issue templates](https://github.com/prequel-dev/cre/issues/new/choose) to help guide you:

* **Bug report**: Create a report to help us improve
* **Feature request**: Suggest an idea for this project
* **New rule**: Suggestions and ideas for new CREs
* **Tune existing rule**: Suggest changes to make to an existing CRE to address false positives or negatives

If none of the templates fit your needs, feel free to open a blank issue and add relevant labels. The more details you provide—such as steps to reproduce, logs, or code snippets—the easier it is for us to help.

### "My issue isn't getting enough attention"

We’re sorry if your issue seems overlooked! We try to respond promptly, but sometimes things get busy. If you feel your issue has been stuck for too long, feel free to:

* Bump the issue with an additional comment.
* Tag a maintainer or reach out to us on Slack (if available).

We value your contributions and want to ensure your feedback is heard.

### "I want to help!"

That’s great! If you’d like to fix a bug or introduce a new rule, please find or open a related issue before you begin coding. That way, we can coordinate efforts, share insights, and avoid duplicating work. Once we agree on the approach, we’ll be happy to collaborate with you to get your changes merged.

## How we use Git and GitHub

### Forking

We follow the [GitHub forking model](https://help.github.com/articles/fork-a-repo/) for collaborating on CRE rules. That means you’ll:

* Fork this repository to your own GitHub account.
* Clone your fork locally.
* Add a remote named upstream pointing to the official CREs repo.

In upcoming sections, you might see references to `upstream`, so be sure you’ve added it properly.

### Commit messages

* Commit often while working on a branch.
* Write helpful, descriptive commit messages (avoid generic messages like “update”).
* Briefly summarize what changed and why in each commit.

(For more tips, check out [this helpful blog post](https://chris.beams.io/posts/git-commit/) on writing good commit messages.)

### What goes into a Pull Request

When submitting a PR, please include:

* A clear description of what you changed and why.
* Links to any relevant issues (e.g., “Closes #123”), external references, or related PRs.
* If adding a new rule:
  * An explanation of how and why it works.
  * Potential impacts on noise (false positives or negatives).
  * Screenshots or other supporting evidence, ensuring you’ve removed any sensitive info.

For more details on our pull request process, see the “Submitting a Pull Request” section below.

## Unique IDs

### CRE IDs

When you open a new PR to add a new rule, you will receive a new CRE identifier in the PR. CRE IDs must be of the form `CRE-YEAR-0123`.

### Rule IDs

Use the `ruler` tool to generate a new rule ID hash.

```bash
$ ./bin/ruler id 
DCejCw6644SvCgdJ5XT3bm
```

You generate more than one rule ID at a time.

```bash
$ ./bin/ruler id -n 5
DXKSK2oKMnPsi8at51bEC6
GMT1B5FUQj7ftGBAiWcFkg
Cz3K1hq64GvVbrvyddUaRQ
S4SRL9GZCGpD33aE3eZPW
T7yUNfneHUjSTLUtb7nq2f
```

Rule ID and hashes have a minimum length requirement of 12 characters. Only alpha-numerics are supported. No special characters.

## Tags and Categories

When adding a new category or tag for a rule, you must also add it to the tags.yaml or categories.yaml file in the `rules/tags/` folder. The name of the tag must be globally unique across all tags and categories. The `description` and `displayName` fields are required.

We try to keep categories more generic and use tags to be more specific. Category should allow bundling rules by problems across specific technologies in tags.

## Schema

It is recommended to add the CRE JSON schema in Cursor or VSCode. Install the [Yaml extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml) and add the following to your settings.json:

```json title="settings.json"
{
    "window.commandCenter": 1,
    "workbench.preferredHighContrastColorTheme": "Visual Studio Dark",
    "yaml.schemas": {
            "https://docs.prequel.dev/cre-schema.json": [
                "/rules/cre-*/*.yaml"
            ]
    },
    "redhat.telemetry.enabled": false,
    "workbench.settings.applyToAllProfiles": [

    ]
}
```

This will enforce the schema on Yaml files in the root of the CRE repository under the `rules/cre-*` directories.

Go to VSCode Settings. Then search for "settings" and click `Edit in settings.json`.

## Building

To be consumed by `preq` or any other problem detector, CREs should be combined into one Yaml document using the latest `ruler` release.

The `ruler` tool also performs the following checks before producing a combined Yaml document:

* Verifies that all referenced tags and categories from CREs exist in the tags.yaml and categories.yaml files
* Verifies there are no duplicate CRE or rule IDs
* Verifies rules do not have any syntax or semantic errors

### From the latest release of `ruler`

Use the `-p` argument to specify the path to the `rules/` folder and `-o` to specify the output folder for the final CRE Yaml document.

```bash
ruler-linux-amd64 build -p rules -o ./bin
Apr 22 16:18:52.941682 INF build.go:208 > Building outPath=./bin vers=v0.3.10
Apr 22 16:18:52.943047 INF build.go:184 > Rule hash=CAbgxyQnLLP12A6GrRHAdcBsbtstGio1gEAj3kLqyRe9 id=CRE-2024-0007
Apr 22 16:18:52.943277 INF build.go:184 > Rule hash=7QuPUNsDLabUWYVKWWbSNeqHZt3vG84VvbnpKqvrA5yr id=CRE-2024-0016
Apr 22 16:18:52.943448 INF build.go:184 > Rule hash=AGYhkt4mrh9j9pmq3BEzmESbjHr8MCNjQa5CRrvmzh2S id=CRE-2024-0043
Apr 22 16:18:52.943659 INF build.go:184 > Rule hash=7DrXpqkkXSCUMKbvmJVEugnpxzVxPTRnjYMvEJoRUi1t id=CRE-2024-0021
Apr 22 16:18:52.943871 INF build.go:184 > Rule hash=3PE3ZmUcEfBkMg3k9StDJmqPAi8FSTPxqAa61YqjpcGw id=CRE-2025-0025
Apr 22 16:18:52.944156 INF build.go:184 > Rule hash=4D5wnjqwGLVgg3DJemdC7uRudZAcMjq22isKiqJ8DoaX id=CRE-2025-0026
Apr 22 16:18:52.944370 INF build.go:184 > Rule hash=Dm7tjHxxRCwdH6YmBe1K26kUp3RNSkCwkE3f1LrdLjDD id=CRE-2025-0027
Apr 22 16:18:52.944620 INF build.go:184 > Rule hash=AQDNkhwFYb9HU5tzUdvmYB84aVUnoxSTEgSA3KP8PJyi id=CRE-2025-0028
Apr 22 16:18:52.944838 INF build.go:184 > Rule hash=39tpJniD6H5K2pvsP29kKm1dmQ3P1JzL7L934aFRu4JL id=CRE-2025-0029
Apr 22 16:18:52.945043 INF build.go:184 > Rule hash=3ZMgEmL3vUYYndeHfUDrDVnETR9hJeU5r3xfsBQfiG6s id=CRE-2025-0030
Apr 22 16:18:52.945250 INF build.go:184 > Rule hash=8U4gFoQrEDjSWTEKZThpYoY7kfGSJ5tdfhoTxXxHY5By id=CRE-2025-0031
Apr 22 16:18:52.945464 INF build.go:184 > Rule hash=44ujRdhMFqc1rGAzEQLPt8MyixRNdnWSupLJMA2SsLf6 id=CRE-2025-0032
Apr 22 16:18:52.945651 INF build.go:184 > Rule hash=Bk7SGsBStG3c8UQW9FSdzP3cYBK5HL3R5HPhZazfyMJe id=CRE-2025-0033
Apr 22 16:18:52.945846 INF build.go:184 > Rule hash=4irDu2Tw3BH8NhGRe7brVW2EkgHu3j7yVTRwJJPc1u5F id=CRE-2025-0034
Apr 22 16:18:52.946057 INF build.go:184 > Rule hash=EzWAm2EvfySF2vZtc3xNAhhY6punkULn5ymZ1wSdwcXA id=CRE-2025-0035
Apr 22 16:18:52.946277 INF build.go:184 > Rule hash=88jxCTk7gozKvFjB4Vpc2ndxfziMANRStrJm79QRUM7W id=CRE-2025-0036
Apr 22 16:18:52.946491 INF build.go:184 > Rule hash=Htzxhggwe3vtCQC3WURCCHTgrU4KSbAhWQYMTWQjgCSQ id=CRE-2025-0037
Apr 22 16:18:52.946674 INF build.go:184 > Rule hash=G6c2SyYC5stJpjAGrG2T9JmZMbX7kbKGVC9bicJX3XHX id=CRE-2025-0038
Apr 22 16:18:52.946909 INF build.go:184 > Rule hash=9mbKubGuV85CQsxkJAszegmQHkiU5qvNMWu1FitdYKGy id=CRE-2025-0039
Apr 22 16:18:52.947124 INF build.go:184 > Rule hash=Dtd1HsWKuzGmVUNoRpTTD9CuXdUgSKa69G2r56jj9v1F id=CRE-2025-0040
Apr 22 16:18:52.947329 INF build.go:184 > Rule hash=3ASBTPgnhNuKUEPfKYyH4sgcKZhZWcCtZ8snQNdqm5iV id=CRE-2025-0041
Apr 22 16:18:52.947535 INF build.go:184 > Rule hash=CNu2xuaY7dKoPXeD9vLYDuTAhrwZLcZhRGGGKpaPs4TV id=CRE-2025-0042
Apr 22 16:18:52.947747 INF build.go:184 > Rule hash=F7rmaDG6dN7ymn66im5HPquFwCK7eyWPM8aFrTRu71bu id=CRE-2025-0043
Wrote file: cre-rules.0.3.10.yaml
```

### From `ruler` source

Ensure you have [Go 1.24.1 or newer](https://go.dev/doc/install) installed.

```bash
$ make
rm -rf bin/*
mkdir -p ./bin
Apr 22 16:04:32.555432 INF build.go:208 > Building outPath=./bin vers=v0.3.10
Apr 22 16:04:32.556809 INF build.go:184 > Rule hash=CAbgxyQnLLP12A6GrRHAdcBsbtstGio1gEAj3kLqyRe9 id=CRE-2024-0007
Apr 22 16:04:32.557037 INF build.go:184 > Rule hash=7QuPUNsDLabUWYVKWWbSNeqHZt3vG84VvbnpKqvrA5yr id=CRE-2024-0016
Apr 22 16:04:32.557216 INF build.go:184 > Rule hash=AGYhkt4mrh9j9pmq3BEzmESbjHr8MCNjQa5CRrvmzh2S id=CRE-2024-0043
Apr 22 16:04:32.557429 INF build.go:184 > Rule hash=7DrXpqkkXSCUMKbvmJVEugnpxzVxPTRnjYMvEJoRUi1t id=CRE-2024-0021
Apr 22 16:04:32.557623 INF build.go:184 > Rule hash=3PE3ZmUcEfBkMg3k9StDJmqPAi8FSTPxqAa61YqjpcGw id=CRE-2025-0025
Apr 22 16:04:32.557900 INF build.go:184 > Rule hash=4D5wnjqwGLVgg3DJemdC7uRudZAcMjq22isKiqJ8DoaX id=CRE-2025-0026
Apr 22 16:04:32.558103 INF build.go:184 > Rule hash=Dm7tjHxxRCwdH6YmBe1K26kUp3RNSkCwkE3f1LrdLjDD id=CRE-2025-0027
Apr 22 16:04:32.558308 INF build.go:184 > Rule hash=AQDNkhwFYb9HU5tzUdvmYB84aVUnoxSTEgSA3KP8PJyi id=CRE-2025-0028
Apr 22 16:04:32.558495 INF build.go:184 > Rule hash=39tpJniD6H5K2pvsP29kKm1dmQ3P1JzL7L934aFRu4JL id=CRE-2025-0029
Apr 22 16:04:32.558691 INF build.go:184 > Rule hash=3ZMgEmL3vUYYndeHfUDrDVnETR9hJeU5r3xfsBQfiG6s id=CRE-2025-0030
Apr 22 16:04:32.558907 INF build.go:184 > Rule hash=8U4gFoQrEDjSWTEKZThpYoY7kfGSJ5tdfhoTxXxHY5By id=CRE-2025-0031
Apr 22 16:04:32.559118 INF build.go:184 > Rule hash=44ujRdhMFqc1rGAzEQLPt8MyixRNdnWSupLJMA2SsLf6 id=CRE-2025-0032
Apr 22 16:04:32.559303 INF build.go:184 > Rule hash=Bk7SGsBStG3c8UQW9FSdzP3cYBK5HL3R5HPhZazfyMJe id=CRE-2025-0033
Apr 22 16:04:32.559487 INF build.go:184 > Rule hash=4irDu2Tw3BH8NhGRe7brVW2EkgHu3j7yVTRwJJPc1u5F id=CRE-2025-0034
Apr 22 16:04:32.559693 INF build.go:184 > Rule hash=EzWAm2EvfySF2vZtc3xNAhhY6punkULn5ymZ1wSdwcXA id=CRE-2025-0035
Apr 22 16:04:32.559920 INF build.go:184 > Rule hash=88jxCTk7gozKvFjB4Vpc2ndxfziMANRStrJm79QRUM7W id=CRE-2025-0036
Apr 22 16:04:32.560158 INF build.go:184 > Rule hash=Htzxhggwe3vtCQC3WURCCHTgrU4KSbAhWQYMTWQjgCSQ id=CRE-2025-0037
Apr 22 16:04:32.560351 INF build.go:184 > Rule hash=G6c2SyYC5stJpjAGrG2T9JmZMbX7kbKGVC9bicJX3XHX id=CRE-2025-0038
Apr 22 16:04:32.560574 INF build.go:184 > Rule hash=9mbKubGuV85CQsxkJAszegmQHkiU5qvNMWu1FitdYKGy id=CRE-2025-0039
Apr 22 16:04:32.560787 INF build.go:184 > Rule hash=Dtd1HsWKuzGmVUNoRpTTD9CuXdUgSKa69G2r56jj9v1F id=CRE-2025-0040
Apr 22 16:04:32.561004 INF build.go:184 > Rule hash=3ASBTPgnhNuKUEPfKYyH4sgcKZhZWcCtZ8snQNdqm5iV id=CRE-2025-0041
Apr 22 16:04:32.561250 INF build.go:184 > Rule hash=CNu2xuaY7dKoPXeD9vLYDuTAhrwZLcZhRGGGKpaPs4TV id=CRE-2025-0042
Apr 22 16:04:32.561468 INF build.go:184 > Rule hash=F7rmaDG6dN7ymn66im5HPquFwCK7eyWPM8aFrTRu71bu id=CRE-2025-0043
Wrote file: cre-rules.0.3.10.yaml
```

## Testing a rule

In the `tests` folder, run:

```bash
cd ./tests
LOG_LEVEL=INFO go test
```

When adding a new rule, you must add a `test.log` file that contains the postive conditions to detect the problem. You are also encouraged to add a `test-fp.log` file that contains the data that would match all of the positive conditions but also contains negative conditions that would indicate a false positive.

## Signing the contributor license agreement

New contributors need to sign the [Contributor License Agreement](https://github.com/prequel-dev/cre/blob/main/cla.md). This ensures that both you and the project are properly protected. Please email contributor-agreement@prequel.dev to complete the CLA process. If you’re unsure whether you need to sign the CLA, feel free to ask on Slack or open an issue.

## Submitting a Pull Request

* Push your local branch to your fork on GitHub.
* Open a Pull Request against `main`.
* Mention the issue you’re addressing, for example, “Closes #123.”

Keeping our conversations linked to relevant issues makes it easier for everyone to follow the discussion and context.

### What to expect from a code review

After submitting your PR:

* A reviewer (possibly you, if you have permissions) will look over your code.
* We’ll likely discuss any changes or ask clarifying questions in the PR comments.
* Even if something works in your environment, we might refine it to ensure it’s broadly applicable.

Don’t be discouraged by feedback; we aim to ensure the highest quality for all users in the problem detection community!

### How we handle merges

To keep our commit history clean and easy to follow, we typically use squash and merge. This means all of your commits will be combined into a single commit in main. If you need the full commit history for your own records, make sure to keep it locally before squashing.

----

Thank you for your contributions, and welcome to the CRE community! If you have any questions or run into issues, feel free to ask on Slack or open a new issue. We’re excited to see what you’ll bring to the project!
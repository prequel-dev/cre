# Contributing to CREs

Thank you for your interest in contributing to Common Reliability Enumerations. We recommend that you read these contribution guidelines carefully. We want to make it easy for you to contribute back to the problem detection community.

These guidelines will also help you post meaningful issues that will be more easily understood, considered, and resolved. These guidelines are here to help you whether you are creating a new rule, opening an issue to report a false positive, or requesting a feature.

## Effective issue creation in CREs

### Why we create issues before contributing code or new rules

We generally create issues in GitHub before contributing code or new rules. This helps front-load the conversation before the rules. There are many rules that will make sense in one or two environments, but don't work as well in general. Some rules are overfitted to a particular indicator or tool. By creating an issue first, it creates an opportunity to bounce our ideas off each other to see what's feasible and what ways to approach detection.

By contrast, starting with a pull request makes it more difficult to revisit the approach. Many PRs are treated as mostly done and shouldn't need much work to get merged. Nobody wants to receive PR feedback that says "start over" or "closing: won't merge." That's discouraging to everyone, and we can avoid those situations if we have the discussion together earlier in the development process. It might be a mental switch for you to start the discussion earlier, but it makes us all more productive and and our rules more effective.

### What a good issue looks like

We have a few types of issue templates to [choose from](https://github.com/prequel-dev/cre/issues/new/choose). If you don't find a template that matches or simply want to ask a question, create a blank issue and add the appropriate labels.

* **Bug report**: Create a report to help us improve
* **Feature request**: Suggest an idea for this project
* **New rule**: Suggestions and ideas for new CREs
* **Tune existing rule**: Suggest changes to make to an existing CRE to address false positives or negatives

When requesting a **New rule**, please create an issue of the **New rule** type. The issue contains a handful of questions about the targeted behavior and the approach to detection:

### "My issue isn't getting enough attention"

First of all, **sorry about that!** We'll tag issues and pull requests with the target release when applicable. Join us on Slack and don't hesitate to reach out to us if you feel like we aren't doing a good job communicating.

Of course, feel free to bump your issues if you think they've been neglected for a prolonged period.

### "I want to help!"

**Great!**. If you have a bug fix or new rule that you would like to contribute, please **find or open an issue about it before you start working on it.** Talk about what you would like to do. It may be that somebody is already working on it, or that there are particular issues that you should know about before implementing the change.

We enjoy working with contributors to get their code accepted. There are many approaches to fixing a problem and it is important to find the best approach before writing too much code.

## How we use Git and GitHub

### Forking

We follow the [GitHub forking model](https://help.github.com/articles/fork-a-repo/) for collaborating on CRE rules. This model assumes that you have a remote called `upstream` which points to the official CRE repo, which we'll refer to in later code snippets.

### Commit messages

* Feel free to make as many commits as you want, while working on a branch.
* Please use your commit messages to include helpful information on your changes. Commit messages that look like `update` are unhelpful to reviewers. Try to be clear and concise with the changes in a commit. Here's a [good blog](https://chris.beams.io/posts/git-commit/) on general best practices for commit messages.

### What goes into a Pull Request

* Please include an explanation of your changes in your PR description.
* Links to relevant issues, external resources, or related PRs are very important and useful.
* Please try to explain *how* and *why* your rule works. Can you explain what makes the logic sound? Does it actually detect what it's supposed to? If you include the screenshot, please make sure to crop out any sensitive information!
* Please try to capture the expectations for noise levels: is the rule prone to false positives or false negatives?
* See [Submitting a Pull Request](#submitting-a-pull-request) for more info.

## Building

### Using the `ruler` build tool

### From source

Ensure you have [Go 1.24.1 or newer](https://go.dev/doc/install) installed.

```bash
$ make
rm -rf bin/*
Apr  6 12:26:22.426820 INF ruler.go:32 > Starting creVersion=0.3.5 gitHash=bab564291a90d398612bb8624f5deb021d396fbf
Apr  6 12:26:22.426952 INF build.go:204 > Building outPath=./bin vers=v0.3.5
Apr  6 12:26:22.427611 INF build.go:180 > Rule hash=3JJigAvM37cTd12UHSUAW62ESCbmsyoP8yaLMG2ciZHn id=CRE-2024-0007
Apr  6 12:26:22.427760 INF build.go:180 > Rule hash=9tYbXspjokxGYy4h77Y22XzMKYKC87cG51rAc5XX6beA id=CRE-2024-0016
Apr  6 12:26:22.427917 INF build.go:180 > Rule hash=BsNNmQfmwjreJdBChDXKCsJbXFerepS4PpCVWEKxdLu1 id=CRE-2024-0021
Wrote file [sha256 b6cea0c37104234650e00807a5ab23096061cd22e3e2d64df74b5358cf97f875]: cre-rules.0.3.5.b6cea0c3.yaml
Wrote hash file: bin/cre-rules.0.3.5.b6cea0c3.yaml.sha256
```

### Testing a rule

## Signing the contributor license agreement

## Submitting a Pull Request

Push your local changes to your forked copy of the repository and submit a Pull Request. In the Pull Request, describe what your changes do and mention the number of the issue where discussion has taken place, e.g., "Closes #123".

Always submit your pull against `main`.

### What to expect from a code review

After a pull is submitted, it needs to get to review. If you have commit permissions on the CRE repo you will probably perform these steps while submitting your Pull Request. If not, a member of the Prequel organization will do them for you, though you can help by suggesting a reviewer for your changes if you've interacted with someone while working on the issue.

Most likely, we will want to have a conversation in the pull request. Please understand that even if a rule is working in your environment, it still may not be a good fit for all users.

### How we handle merges

We recognize that Git commit messages are a history of all changes to the repository. We want to make this history easy to read and as concise and clear as possible. When we merge a pull request, we squash commits using GitHub's "Squash and Merge" method of merging. This keeps a clear history to the repository, since we rarely need to know about the commits that happen *within* a working branch for a pull request.

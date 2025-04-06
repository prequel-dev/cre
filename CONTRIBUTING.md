# Contributing to CREs

Thank you for your interest in contributing to Common Reliability Enumerations. We recommend that you read these contribution guidelines carefully. We want to make it easy for you to contribute back to the problem detection community.

These guidelines will also help you post meaningful issues that will be more easily understood, considered, and resolved. These guidelines are here to help you whether you are creating a new rule, opening an issue to report a false positive, or requesting a feature.

## Table of Contents

## Effective issue creation in CREs

### Why we create issues before contributing code or new rules

We generally create issues in GitHub before contributing code or new rules. This helps front-load the conversation before the rules. There are many rules that will make sense in one or two environments, but don't work as well in general. Some rules are overfitted to a particular indicator or tool. By creating an issue first, it creates an opportunity to bounce our ideas off each other to see what's feasible and what ways to approach detection.

By contrast, starting with a pull request makes it more difficult to revisit the approach. Many PRs are treated as mostly done and shouldn't need much work to get merged. Nobody wants to receive PR feedback that says "start over" or "closing: won't merge." That's discouraging to everyone, and we can avoid those situations if we have the discussion together earlier in the development process. It might be a mental switch for you to start the discussion earlier, but it makes us all more productive and and our rules more effective.

### What a good issue looks like

We have a few types of issue templates to [choose from](https://github.com/elastic/detection-rules/issues/new/choose). If you don't find a template that matches or simply want to ask a question, create a blank issue and add the appropriate labels.

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

### Testing a rule with the CLI

When a rule is ready, it can be tested with unit tests. Detection Rules has several tests that run locally to validate rules in the repository. These tests make sure that rules are syntactically correct, use ECS or Beats schemas correctly, and ensure that metadata is also validated. There are also internal tests to make sure that the tools and functions to manager the repository are working as expected.

To run tests, simply run the command `test` with the CLI
```console
$ python -m detection_rules test

============================================================= test session starts ==============================================================
collected 73 items

tests/test_all_rules.py::TestValidRules::test_all_rule_files PASSED                                                                                               [  1%]
tests/test_all_rules.py::TestValidRules::test_all_rule_queries_optimized PASSED                                                                                   [  2%]
tests/test_all_rules.py::TestValidRules::test_all_rules_as_rule_schema PASSED                                                                                     [  4%]
tests/test_all_rules.py::TestValidRules::test_all_rules_tuned PASSED                                                                                              [  5%]
...
tests/kuery/test_parser.py::ParserTests::test_number_exists PASSED                                                                                                [ 98%]
tests/kuery/test_parser.py::ParserTests::test_number_wildcard_fail PASSED                                                                                         [100%]

========================================================================== 73 passed in 45.47s ==========================================================================
```


## Writing style

Our rules are much more than queries. We capture a lot of metadata within the rules, such as severity, index pattterns, and noise level. We also have several fields that are user-readable text, such as `name`, `description`, `false_positives`, `investigation_notes`, and `name`. Those fields, which are populated with English text[*](#i18n-note), should follow the [Elastic UI writing guidelines](https://elastic.github.io/eui/#/guidelines/writing). We want our text to be *clear* and *concise*, *consistent* and *conversational*.

<a id="i18n-note">\* Note</a>: We currently don't have i18n support for Detection Rules.


## Signing the contributor license agreement

Please make sure you've signed the [Contributor License Agreement](http://www.elastic.co/contributor-agreement/). We're not asking you to assign copyright to us, but to give us the right to distribute your code without restriction. We ask this of all contributors in order to assure our users of the origin and continuing existence of the code. You only need to sign the CLA once.


## Submitting a Pull Request

Push your local changes to your forked copy of the repository and submit a Pull Request. In the Pull Request, describe what your changes do and mention the number of the issue where discussion has taken place, e.g., "Closes #123".

Always submit your pull against `main` unless you are making changes for the pending release during feature freeze (see [Branching](#branching) for our branching strategy).

Then sit back and wait. We will probably have a discussion in the pull request and may request changes before merging. We're not trying to get in the way, but want to work with you to get your contributions in Detection Rules.


### What to expect from a code review

After a pull is submitted, it needs to get to review. If you have commit permissions on the Detection Rules repo you will probably perform these steps while submitting your Pull Request. If not, a member of the Elastic organization will do them for you, though you can help by suggesting a reviewer for your changes if you've interacted with someone while working on the issue.

Most likely, we will want to have a conversation in the pull request. We want to encourage contributions, but we also want to keep in mind how changes may affect other Elastic users. Please understand that even if a rule is working in your environment, it still may not be a good fit for all users of Elastic Security.

### How we handle merges

We recognize that Git commit messages are a history of all changes to the repository. We want to make this history easy to read and as concise and clear as possible. When we merge a pull request, we squash commits using GitHub's "Squash and Merge" method of merging. This keeps a clear history to the repository, since we rarely need to know about the commits that happen *within* a working branch for a pull request.

The exception to this rule is backport PRs. We want to maintain that commit history, because the commits within a release branch have already been squashed. If we were to squash again to a single commit, we would just see a commit "Backport changes from `{majorVersion.minorVersion}`" show up in main. This would obscure the changes. For backport pull requests, we will either "Create a Merge Commit" or "Rebase and Merge." For more information, see [Branching](#branching) for our branching strategy.

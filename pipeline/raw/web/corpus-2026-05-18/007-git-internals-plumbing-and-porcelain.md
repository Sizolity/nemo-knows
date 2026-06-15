---
title: Git Internals - Plumbing and Porcelain
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain
tags: [git, web-corpus]
confidence: medium
---

# Git Internals - Plumbing and Porcelain

## Fetch Metadata

- Corpus item: 7
- Category: Git
- Source URL: https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain
- Final URL: https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Deep implementation source for concept pages.
- Fetched page title: Git - Plumbing and Porcelain

## Retrieved Text

Git
____________________
[dark-mode.svg]

* About
+ Trademark
* Learn
+ Book
+ Cheat Sheet
+ Videos
+ External Links
* Tools
+ Command Line
+ GUIs
+ Hosting
* Reference
* Install
* Community
__________________________________________________________________

This book is available in English.

Full translation available in
azərbaycan dili,
български език,
Deutsch,
Español,
فارسی,
Français,
Ελληνικά,
日本語,
한국어,
Nederlands,
Русский,
Slovenščina,
Српски,
Svenska,
Tagalog,
Türkçe.
Українська,
简体中文,

Partial translations available in
Čeština,
Македонски,
Polski,
Português (Brasil),
Ўзбекча,
繁體中文,

Translations started for
Беларуская,
Indonesian,
Italiano,
Bahasa Melayu,
Português (Portugal).
__________________________________________________________________

The source of this book is hosted on GitHub.
Patches, suggestions and comments are welcome.
Chapters ▾
1. 1. Getting Started
1. 1.1 About Version Control
2. 1.2 A Short History of Git
3. 1.3 What is Git?
4. 1.4 The Command Line
5. 1.5 Installing Git
6. 1.6 First-Time Git Setup
7. 1.7 Getting Help
8. 1.8 Summary
2. 2. Git Basics
1. 2.1 Getting a Git Repository
2. 2.2 Recording Changes to the Repository
3. 2.3 Viewing the Commit History
4. 2.4 Undoing Things
5. 2.5 Working with Remotes
6. 2.6 Tagging
7. 2.7 Git Aliases
8. 2.8 Summary
3. 3. Git Branching
1. 3.1 Branches in a Nutshell
2. 3.2 Basic Branching and Merging
3. 3.3 Branch Management
4. 3.4 Branching Workflows
5. 3.5 Remote Branches
6. 3.6 Rebasing
7. 3.7 Summary
4. 4. Git on the Server
1. 4.1 The Protocols
2. 4.2 Getting Git on a Server
3. 4.3 Generating Your SSH Public Key
4. 4.4 Setting Up the Server
5. 4.5 Git Daemon
6. 4.6 Smart HTTP
7. 4.7 GitWeb
8. 4.8 GitLab
9. 4.9 Third Party Hosted Options
10. 4.10 Summary
5. 5. Distributed Git
1. 5.1 Distributed Workflows
2. 5.2 Contributing to a Project
3. 5.3 Maintaining a Project
4. 5.4 Summary

1. 6. GitHub
1. 6.1 Account Setup and Configuration
2. 6.2 Contributing to a Project
3. 6.3 Maintaining a Project
4. 6.4 Managing an organization
5. 6.5 Scripting GitHub
6. 6.6 Summary
2. 7. Git Tools
1. 7.1 Revision Selection
2. 7.2 Interactive Staging
3. 7.3 Stashing and Cleaning
4. 7.4 Signing Your Work
5. 7.5 Searching
6. 7.6 Rewriting History
7. 7.7 Reset Demystified
8. 7.8 Advanced Merging
9. 7.9 Rerere
10. 7.10 Debugging with Git
11. 7.11 Submodules
12. 7.12 Bundling
13. 7.13 Replace
14. 7.14 Credential Storage
15. 7.15 Summary
3. 8. Customizing Git
1. 8.1 Git Configuration
2. 8.2 Git Attributes
3. 8.3 Git Hooks
4. 8.4 An Example Git-Enforced Policy
5. 8.5 Summary
4. 9. Git and Other Systems
1. 9.1 Git as a Client
2. 9.2 Migrating to Git
3. 9.3 Summary
5. 10. Git Internals
1. 10.1 Plumbing and Porcelain
2. 10.2 Git Objects
3. 10.3 Git References
4. 10.4 Packfiles
5. 10.5 The Refspec
6. 10.6 Transfer Protocols
7. 10.7 Maintenance and Data Recovery
8. 10.8 Environment Variables
9. 10.9 Summary

1. A1. Appendix A: Git in Other Environments
1. A1.1 Graphical Interfaces
2. A1.2 Git in Visual Studio
3. A1.3 Git in Visual Studio Code
4. A1.4 Git in IntelliJ / PyCharm / WebStorm / PhpStorm /
RubyMine
5. A1.5 Git in Sublime Text
6. A1.6 Git in Bash
7. A1.7 Git in Zsh
8. A1.8 Git in PowerShell
9. A1.9 Summary
2. A2. Appendix B: Embedding Git in your Applications
1. A2.1 Command-line Git
2. A2.2 Libgit2
3. A2.3 JGit
4. A2.4 go-git
5. A2.5 Dulwich
3. A3. Appendix C: Git Commands
1. A3.1 Setup and Config
2. A3.2 Getting and Creating Projects
3. A3.3 Basic Snapshotting
4. A3.4 Branching and Merging
5. A3.5 Sharing and Updating Projects
6. A3.6 Inspection and Comparison
7. A3.7 Debugging
8. A3.8 Patching
9. A3.9 Email
10. A3.10 External Systems
11. A3.11 Administration
12. A3.12 Plumbing Commands

2nd Edition

10.1 Git Internals - Plumbing and Porcelain

You may have skipped to this chapter from a much earlier chapter, or
you may have gotten here after sequentially reading the entire book up
to this point — in either case, this is where we’ll go over the inner
workings and implementation of Git. We found that understanding this
information was fundamentally important to appreciating how useful and
powerful Git is, but others have argued to us that it can be confusing
and unnecessarily complex for beginners. Thus, we’ve made this
discussion the last chapter in the book so you could read it early or
later in your learning process. We leave it up to you to decide.

Now that you’re here, let’s get started. First, if it isn’t yet clear,
Git is fundamentally a content-addressable filesystem with a VCS user
interface written on top of it. You’ll learn more about what this means
in a bit.

In the early days of Git (mostly pre 1.5), the user interface was much
more complex because it emphasized this filesystem rather than a
polished VCS. In the last few years, the UI has been refined until it’s
as clean and easy to use as any system out there; however, the
stereotype lingers about the early Git UI that was complex and
difficult to learn.

The content-addressable filesystem layer is amazingly cool, so we’ll
cover that first in this chapter; then, you’ll learn about the
transport mechanisms and the repository maintenance tasks that you may
eventually have to deal with.

Plumbing and Porcelain

This book covers primarily how to use Git with 30 or so subcommands
such as checkout, branch, remote, and so on. But because Git was
initially a toolkit for a version control system rather than a full
user-friendly VCS, it has a number of subcommands that do low-level
work and were designed to be chained together UNIX-style or called from
scripts. These commands are generally referred to as Git’s “plumbing”
commands, while the more user-friendly commands are called “porcelain”
commands.

As you will have noticed by now, this book’s first nine chapters deal
almost exclusively with porcelain commands. But in this chapter, you’ll
be dealing mostly with the lower-level plumbing commands, because they
give you access to the inner workings of Git, and help demonstrate how
and why Git does what it does. Many of these commands aren’t meant to
be used manually on the command line, but rather to be used as building
blocks for new tools and custom scripts.

When you run git init in a new or existing directory, Git creates the
.git directory, which is where almost everything that Git stores and
manipulates is located. If you want to back up or clone your
repository, copying this single directory elsewhere gives you nearly
everything you need. This entire chapter basically deals with what you
can see in this directory. Here’s what a newly-initialized .git
directory typically looks like:
$ ls -F1
config
description
HEAD
hooks/
info/
objects/
refs/

Depending on your version of Git, you may see some additional content
there, but this is a fresh git init repository — it’s what you see by
default. The description file is used only by the GitWeb program, so
don’t worry about it. The config file contains your project-specific
configuration options, and the info directory keeps a global exclude
file for ignored patterns that you don’t want to track in a .gitignore
file. The hooks directory contains your client- or server-side hook
scripts, which are discussed in detail in Git Hooks.

This leaves four important entries: the HEAD and (yet to be created)
index files, and the objects and refs directories. These are the core
parts of Git. The objects directory stores all the content for your
database, the refs directory stores pointers into commit objects in
that data (branches, tags, remotes and more), the HEAD file points to
the branch you currently have checked out, and the index file is where
Git stores your staging area information. You’ll now look at each of
these sections in detail to see how Git operates.
prev | next
About this site
Patches, suggestions, and comments are welcome.
Git is a member of Software Freedom Conservancy

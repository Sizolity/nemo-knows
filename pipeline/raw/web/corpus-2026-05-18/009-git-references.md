---
title: Git References
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://git-scm.com/book/en/v2/Git-Internals-Git-References
tags: [git, web-corpus]
confidence: medium
---

# Git References

## Fetch Metadata

- Corpus item: 9
- Category: Git
- Source URL: https://git-scm.com/book/en/v2/Git-Internals-Git-References
- Final URL: https://git-scm.com/book/en/v2/Git-Internals-Git-References
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Good for linking branches, refs, HEAD, and commits.
- Fetched page title: Git - Git References

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

10.3 Git Internals - Git References

Git References

If you were interested in seeing the history of your repository
reachable from commit, say, 1a410e, you could run something like git
log 1a410e to display that history, but you would still have to
remember that 1a410e is the commit you want to use as the starting
point for that history. Instead, it would be easier if you had a file
in which you could store that SHA-1 value under a simple name so you
could use that simple name rather than the raw SHA-1 value.

In Git, these simple names are called “references” or “refs”; you can
find the files that contain those SHA-1 values in the .git/refs
directory. In the current project, this directory contains no files,
but it does contain a simple structure:
$ find .git/refs
.git/refs
.git/refs/heads
.git/refs/tags
$ find .git/refs -type f

To create a new reference that will help you remember where your latest
commit is, you can technically do something as simple as this:
$ echo 1a410efbd13591db07496601ebc7a059dd55cfe9 > .git/refs/heads/master

Now, you can use the head reference you just created instead of the
SHA-1 value in your Git commands:
$ git log --pretty=oneline master
1a410efbd13591db07496601ebc7a059dd55cfe9 Third commit
cac0cab538b970a37ea1e769cbbde608743bc96d Second commit
fdf4fc3344e67ab068f836878b6c4951e3b15f3d First commit

You aren’t encouraged to directly edit the reference files; instead,
Git provides the safer command git update-ref to do this if you want to
update a reference:
$ git update-ref refs/heads/master 1a410efbd13591db07496601ebc7a059dd55cfe9

That’s basically what a branch in Git is: a simple pointer or reference
to the head of a line of work. To create a branch back at the second
commit, you can do this:
$ git update-ref refs/heads/test cac0ca

Your branch will contain only work from that commit down:
$ git log --pretty=oneline test
cac0cab538b970a37ea1e769cbbde608743bc96d Second commit
fdf4fc3344e67ab068f836878b6c4951e3b15f3d First commit

Now, your Git database conceptually looks something like this:
Git directory objects with branch head references included
Figure 176. Git directory objects with branch head references included

When you run commands like git branch <branch>, Git basically runs that
update-ref command to add the SHA-1 of the last commit of the branch
you’re on into whatever new reference you want to create.

The HEAD

The question now is, when you run git branch <branch>, how does Git
know the SHA-1 of the last commit? The answer is the HEAD file.

Usually the HEAD file is a symbolic reference to the branch you’re
currently on. By symbolic reference, we mean that unlike a normal
reference, it contains a pointer to another reference.

However in some rare cases the HEAD file may contain the SHA-1 value of
a Git object. This happens when you checkout a tag, commit, or remote
branch, which puts your repository in "detached HEAD" state.

If you look at the file, you’ll normally see something like this:
$ cat .git/HEAD
ref: refs/heads/master

If you run git checkout test, Git updates the file to look like this:
$ cat .git/HEAD
ref: refs/heads/test

When you run git commit, it creates the commit object, specifying the
parent of that commit object to be whatever SHA-1 value the reference
in HEAD points to.

You can also manually edit this file, but again a safer command exists
to do so: git symbolic-ref. You can read the value of your HEAD via
this command:
$ git symbolic-ref HEAD
refs/heads/master

You can also set the value of HEAD using the same command:
$ git symbolic-ref HEAD refs/heads/test
$ cat .git/HEAD
ref: refs/heads/test

You can’t set a symbolic reference outside of the refs style:
$ git symbolic-ref HEAD test
fatal: Refusing to point HEAD outside of refs/

Tags

We just finished discussing Git’s three main object types (blobs, trees
and commits), but there is a fourth. The tag object is very much like a
commit object — it contains a tagger, a date, a message, and a pointer.
The main difference is that a tag object generally points to a commit
rather than a tree. It’s like a branch reference, but it never
moves — it always points to the same commit but gives it a friendlier
name.

As discussed in Git Basics, there are two types of tags: annotated and
lightweight. You can make a lightweight tag by running something like
this:
$ git update-ref refs/tags/v1.0 cac0cab538b970a37ea1e769cbbde608743bc96d

That is all a lightweight tag is — a reference that never moves. An
annotated tag is more complex, however. If you create an annotated tag,
Git creates a tag object and then writes a reference to point to it
rather than directly to the commit. You can see this by creating an
annotated tag (using the -a option):
$ git tag -a v1.1 1a410efbd13591db07496601ebc7a059dd55cfe9 -m 'Test tag'

Here’s the object SHA-1 value it created:
$ cat .git/refs/tags/v1.1
9585191f37f7b0fb9444f35a9bf50de191beadc2

Now, run git cat-file -p on that SHA-1 value:
$ git cat-file -p 9585191f37f7b0fb9444f35a9bf50de191beadc2
object 1a410efbd13591db07496601ebc7a059dd55cfe9
type commit
tag v1.1
tagger Scott Chacon <schacon@gmail.com> Sat May 23 16:48:58 2009 -0700

Test tag

Notice that the object entry points to the commit SHA-1 value that you
tagged. Also notice that it doesn’t need to point to a commit; you can
tag any Git object. In the Git source code, for example, the maintainer
has added their GPG public key as a blob object and then tagged it. You
can view the public key by running this in a clone of the Git
repository:
$ git cat-file blob junio-gpg-pub

The Linux kernel repository also has a non-commit-pointing tag
object — the first tag created points to the initial tree of the import
of the source code.

Remotes

The third type of reference that you’ll see is a remote reference. If
you add a remote and push to it, Git stores the value you last pushed
to that remote for each branch in the refs/remotes directory. For
instance, you can add a remote called origin and push your master
branch to it:
$ git remote add origin git@github.com:schacon/simplegit-progit.git
$ git push origin master
Counting objects: 11, done.
Compressing objects: 100% (5/5), done.
Writing objects: 100% (7/7), 716 bytes, done.
Total 7 (delta 2), reused 4 (delta 1)
To git@github.com:schacon/simplegit-progit.git
a11bef0..ca82a6d master -> master

Then, you can see what the master branch on the origin remote was the
last time you communicated with the server, by checking the
refs/remotes/origin/master file:
$ cat .git/refs/remotes/origin/master
ca82a6dff817ec66f44342007202690a93763949

Remote references differ from branches (refs/heads references) mainly
in that they’re considered read-only. You can git checkout to one, but
Git won’t symbolically reference HEAD to one, so you’ll never update it
with a commit command. Git manages them as bookmarks to the last known
state of where those branches were on those servers.
prev | next
About this site
Patches, suggestions, and comments are welcome.
Git is a member of Software Freedom Conservancy

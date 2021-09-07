---
title: "Python Buildpack Reference"
menu:
  main:
    parent: reference
    identifier: python-reference
    name: "Python Buildpack"
---

This reference documentation offers an in-depth description of the behavior and
configuration options of the [Paketo Python
Buildpack](https://github.com/paketo-buildpacks/python).  For explanations of
how to use the buildpack for several common use-cases, see the Python How To
[documentation](/docs/howto/python).

## Supported Dependencies
The Python buildpack supports several versions of CPython, Pip and Pipenv.  For
more details on the specific versions supported in a given buildpack version,
see the [release notes](https://github.com/paketo-buildpacks/python/releases).

## <a id="environment-variables"></a> Buildpack-Set Environment Variables

The Python Buildpack sets a few environment variables during the `build` and
`launch` phases of the app lifecycle. The sections below describe each
environment variable and its impact on your app.

### <a id="environment-variable-pythonpath"></a> PYTHONPATH

The [`PYTHONPATH`](https://docs.python.org/3/using/cmdline.html#envvar-PYTHONPATH)
environment variable is used to add directories where python will look for
modules.

* Set by: `CPython`, `Pip` and `Pipenv`
* Phases: `build` and `launch`

The CPython buildpack sets the `PYTHONPATH` value to its installation location,
and the Pip, Pipenv buildpack prepends their `site-packages` location to it.
`site-packages` is the target directory where packages are installed to.

### <a id="environment-variable-pythonuserbase"></a> PYTHONUSERBASE

The [`PYTHONUSERBASE`](https://docs.python.org/3/using/cmdline.html#envvar-PYTHONUSERBASE)
environment variable is used to set the user base directory.

* Set by: `Pip Install` and `Pipenv Install`
* Phases: `build` and `launch`

The value of `PYTHONUSERBASE` is set to the location where these buildapcks install
the application packages so that it can be consumed by the app source code.

### <a id="start-command"></a> Start Command

The Python Buildpack sets the default start command `python`. This starts the Python
REPL (read-eval-print loop) at launch.

The Python Buildpack comes with support for
[`Procfile`](https://paketo.io/docs/buildpacks/configuration/#procfiles)
that lets users set custom start commands easily.

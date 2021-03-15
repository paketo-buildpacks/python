# Restructure to improve composition

## Proposal

The existing Python language buildpacks will be broken down into a set of
buildpacks that, by themselves, provide more limited functionality, but in
composition, fill the same feature set that the language family already
provides.

The current set of buildpacks will be reorganized into the following:

![Proposed Structure](/rfcs/assets/0001-proposed.png)

* `cpython`
  * provides: `cpython`
  * requires: none

* `pip`
  * provides: `pip`
  * requires: `cpython` during `build`

* `pip-install`
  * provides: `site-packages`
  * requires: `cpython` and `pip` during `build`

* `pipenv`
  * provides: `pipenv`
  * requires: `cpython` and `pip` during `build`

* `pipenv-install`
  * provides: `site-packages`
  * requires: `cpython` and `pipenv` during `build`

* `miniconda`
  * provides: `conda`
  * requires: none

* `conda-env-update`
  * provides: `conda-environment`
  * requires: `conda` during `build`

* `python-start`
  * provides: none
  * requires: `cpython` during `launch` OR {`cpython`, `site-packages`} during `launch` OR  `conda-environment` during `launch`

The above implementation buildpacks will be structured into the following order groupings:

```
[[order]]

  [[order.group]]
    id = "paketo-community/cpython"

  [[order.group]]
    id = "paketo-community/pip"

  [[order.group]]
    id = "paketo-community/pipenv"

  [[order.group]]
    id = "paketo-community/pipenv-install"

  [[order.group]]
    id = "paketo-community/python-start"

  [[order.group]]
    id = "paketo-buildpacks/procfile"
    optional = true

[[order]]

  [[order.group]]
    id = "paketo-community/cpython"

  [[order.group]]
    id = "paketo-community/pip"

  [[order.group]]
    id = "paketo-community/pip-install"

  [[order.group]]
    id = "paketo-community/python-start"

  [[order.group]]
    id = "paketo-buildpacks/procfile"
    optional = true

[[order]]

  [[order.group]]
    id = "paketo-community/miniconda"

  [[order.group]]
    id = "paketo-community/conda-env-update"

  [[order.group]]
    id = "paketo-community/python-start"

  [[order.group]]
    id = "paketo-buildpacks/procfile"
    optional = true

[[order]]

  [[order.group]]
    id = "paketo-community/cpython"

  [[order.group]]
    id = "paketo-community/python-start"

  [[order.group]]
    id = "paketo-buildpacks/procfile"
    optional = true
```

## Motivation

This restructuring will allow us to reuse these buildpacks in new features
without needing to reimplement existing code. We can reuse existing ecosystem
tooling like `pip` to aid in the installation of packages using `pipenv`.
Additionally, reducing the responsibilities of each buildpack will help to make
them easier to understand and maintain. We will end up with more buildpacks,
with simpler implementations.

## Implementation

### Package Management

#### Split package manager dependencies from package management processes

The `pip` and `pipenv` buildpacks will only install their respective
dependencies. The current `pip` and `pipenv` buildpacks rely on bootstrapping
their installations by running some form of `python -m pip install...`.

The `pipenv-install` buildpack will be responsible for calling `pipenv install`
directly, instead of delegating to `pip install`. This falls more inline with
what users expect. The existng `pipenv` buildpack was transforming the
`Pipfile` into a `requirements.txt` file and then we'd run the `pip` buildpack
build process as part of that group.

The new implementation creates separate layers for pip, pipenv, and python to
install app dependencies, eliminating the need for the workaround described
above. This involves a few configurations.

The `pipenv install` command will make use of two flags:
* The `--deploy` flag to ensure the Pipfile.lock is up to date, and
* The `--system` flag to make sure Pipfile dependencies are installed into the
  parent system.

The installation commands will set and make use of a few environment variables:
* `PIP_USER` set to 1 to tell pip to install packages to `PYTHONUSERBASE`
* `PYTHONUSERBASE` specifies where packages will be installed to (layer dir)
* `PYTHONPATH` works like the `PATH` variable for looking up Python packages.
  Once packages are installed to `PYTHONUSERBASE`, they should be included on
  the `PYTHONPATH`.

For example if `PYTHONUSERBASE` is set to `/tmp`, when `pip install --user
SomePackage` is run then packages will be found
at`/tmp/lib/python<version>/site-packages/SomePackage`. The python runtime will
locate`SomePackage` via the `PYTHONPATH` variable.

#### Split conda buildpack

The current `conda` buildpack will be divided into separate buildpacks for each
of its constituent processes.  Thus, the `miniconda` buildpack will download
the `Miniconda3` artifact and run its install script to provide `conda`. The
`conda-env-update` buildpack requires `conda` and will run `conda env update`
to update the created environment based on the environment file in app source
code.

### Start Commands

#### Basic Python start command

The `python-start` buildpack will deterministically set `python` as the default
start command, which will start the Python REPL (read-eval-print-loop) at
launch. This will look like: `docker run -it <app-image> python`

From there, users will be able to run a custom start command according
to their needs and the structure of their app's source code. `python-start`
also detects the app's dependencies and requires them at launch.

#### Remove duplicative `Procfile` support

`Procfile` support will be factored out to be handled by the [`Procfile`
buildpack](https://github.com/paketo-buildpacks/procfile) instead of living
directly inside the Python buildpacks, as several of the current buildpacks do.
This change along with the inclusion of the generic `python-start` buildpack
means that users are no longer required to include a Procfile in their source
code for their app image to be usable, though they may still choose to do so.

## Source Material

* https://docs.python.org/3/using/cmdline.html#environment-variables

**Revision History**

* (03/15/2021) Edit: Fix pip-env order grouping - pip was incorrectly removed
  from the order group.

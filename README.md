# Python Paketo Buildpack

## `gcr.io/paketo-buildpacks/python`

The Python Paketo Buildpack provides a set of collaborating buildpacks to build a Python-based application
These buildpacks include:
- [CPython CNB](https://github.com/paketo-buildpacks/cpython)
- [Pipenv CNB](https://github.com/paketo-buildpacks/pipenv)
- [Pipenv Install CNB](https://github.com/paketo-buildpacks/pipenv-install)
- [Pip CNB](https://github.com/paketo-buildpacks/pip)
- [Pip Install CNB](https://github.com/paketo-buildpacks/pip-install)
- [Miniconda CNB](https://github.com/paketo-buildpacks/miniconda)
- [Conda Env Update CNB](https://github.com/paketo-buildpacks/conda-env-update)
- [Poetry CNB](https://github.com/paketo-buildpacks/poetry)
- [Poetry Install CNB](https://github.com/paketo-buildpacks/poetry-install)
- [Poetry Run CNB](https://github.com/paketo-buildpacks/poetry-run)
- [Python Start CNB](https://github.com/paketo-buildpacks/python-start)

Additionally, the following utility buildpacks are included for all application types
- [CA Certificates CNB](https://github.com/paketo-buildpacks/ca-certificates)
- [Watchexec CNB](https://github.com/paketo-buildpacks/watchexec)
- [Procfile CNB](https://github.com/paketo-buildpacks/procfile)
- [Environment Variables CNB](https://github.com/paketo-buildpacks/environment-variables)
- [Image Labels CNB](https://github.com/paketo-buildpacks/image-labels)

The buildpack supports building simple Python applications or applications which
utilize either [Conda](https://conda.io),
[Pipenv](https://pypi.org/project/pipenv/),
[Pip](https://pip.pypa.io/),
or [Poetry](https://python-poetry.org/) for managing their dependencies.

Check out the [Python Paketo Buildpack docs](https://paketo.io/docs/howto/python/) for sample usage and more information.

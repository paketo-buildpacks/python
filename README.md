# Python Paketo Buildpack

## `gcr.io/paketo-community/python`

The Python Paketo Buildpack provides a set of collaborating buildpacks that
enable the building of a Python-based application. These buildpacks include:
- [CPython CNB](https://github.com/paketo-community/cpython)
- [Pipenv CNB](https://github.com/paketo-community/pipenv)
- [Pipenv Install CNB](https://github.com/paketo-community/pipenv-install)
- [Pip CNB](https://github.com/paketo-community/pip)
- [Pip Install CNB](https://github.com/paketo-community/pip-install)
- [Miniconda CNB](https://github.com/paketo-community/miniconda)
- [Conda Env Update CNB](https://github.com/paketo-community/conda-env-update)
- [Python Start CNB](https://github.com/paketo-community/python-start)
- [Procfile CNB](https://github.com/paketo-buildpacks/procfile)

The buildpack supports building simple Python applications or applications which
utilize either [Conda](https://conda.io),
[Pipenv](https://pypi.org/project/pipenv/) or [Pip](https://pip.pypa.io/) for
managing their dependencies.

# Ruby Cloud Native Buildpack

The Python Cloud Native Buildpack provides a set of collaborating buildpacks that
enable the building of a Python-based application. These buildpacks include:
- [Python Runtime CNB](https://github.com/paketo-community/python-runtime)
- [Pipenv CNB](https://github.com/paketo-community/pipenv)
- [Pip CNB](https://github.com/paketo-community/pip)
- [Conda CNB](https://github.com/paketo-community/conda)

The buildpack supports building simple Python applications or applications which
utilize either [Conda](https://conda.io) or [Pipenv](https://pypi.org/project/pipenv/) for managing their dependencies.

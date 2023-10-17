#!/usr/local/sal/Python.framework/Versions/Current/bin/python3

import json
import pathlib

import sal


TCC_LOG = pathlib.Path('/usr/local/sal/tcc_results.json')


def main():
    results = {}
    if TCC_LOG.exists():
        try:
            results = json.loads(TCC_LOG.read_text())
        except ValueError:
            pass
    for entry in results:
        sal.set_checkin_results(entry, results[entry])


if __name__ == "__main__":
    main()

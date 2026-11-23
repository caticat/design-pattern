import io
import json
import unittest
from contextlib import redirect_stdout
from unittest.mock import patch

import playground


class PlaygroundRunTest(unittest.TestCase):
    def test_success_returns_output(self):
        response = {"Status": 0, "Events": [{"Message": "ok\n"}]}
        with patch.object(playground.urllib.request, "urlopen", return_value=self.response(response)):
            self.assertEqual(playground.run(b"package main"), "ok\n")

    def test_nonzero_exit_is_not_shared(self):
        response = {"Status": 2, "Events": [{"Message": "panic: test\n"}]}
        log = io.StringIO()
        with patch.object(playground.urllib.request, "urlopen", return_value=self.response(response)) as request:
            with redirect_stdout(log):
                self.assertIsNone(playground.run(b"package main"))
        request.assert_called_once()
        self.assertIn("panic: test", log.getvalue())

    @staticmethod
    def response(data):
        return io.BytesIO(json.dumps(data).encode("utf-8"))

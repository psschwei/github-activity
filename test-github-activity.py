# Assisted by watsonx Code Assistant 
import os
import pytest
import requests

def test_run_github_query():
    username = "example_user"
    start_dt = "2022-01-01"
    end_dt = "2022-12-31"
    response = run_github_query(username, start_dt, end_dt)
    assert response["data"]["user"]["contributionsCollection"] is not None

def test_run_github_query_invalid_username():
    username = "invalid_user"
    start_dt = "2022-01-01"
    end_dt = "2022-12-31"
    response = run_github_query(username, start_dt, end_dt)
    assert response["data"]["user"]["contributionsCollection"] is None

def test_run_github_query_invalid_dates():
    username = "example_user"
    start_dt = "2022-13-01"
    end_dt = "2022-01-31"
    response = run_github_query(username, start_dt, end_dt)
    assert response["data"]["user"]["contributionsCollection"] is None

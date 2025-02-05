# processing/activities_python.py

import random
from temporalio import activity

@activity.defn(name="PythonAddRandomPrefixActivity")
async def python_add_random_prefix_activity(data: str) -> str:
    prefixes = [
        "alpha-", "beta-", "gamma-", "delta-", "epsilon-", "zeta-",
        "eta-", "theta-", "iota-", "kappa-", "lambda-", "mu-", "nu-",
        "xi-", "omicron-", "pi-", "rho-", "sigma-", "tau-", "upsilon-",
        "phi-", "chi-", "psi-", "omega-"
    ]
    prefix = random.choice(prefixes)
    return f"{prefix}{data}"

import inflect

def pluralize(word):
    """Pluralize a word using the inflect library."""
    engine = inflect.engine()
    return engine.plural(word)

# A function to add filters to the Jinja2 environment
def setup(env):
    """Add custom filters to the Jinja2 environment."""
    env.filters['pluralize'] = pluralize
    return env

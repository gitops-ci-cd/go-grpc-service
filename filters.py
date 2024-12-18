import inflect

def pluralize(word):
    """Pluralize a word using the inflect library."""
    engine = inflect.engine()
    return engine.plural(word)

class CustomExtension:
    def __init__(self, environment):
        """Add custom filters to the Jinja2 environment."""
        environment.filters['pluralize'] = pluralize

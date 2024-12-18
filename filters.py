import inflect
from jinja2.ext import Extension

def pluralize(word):
    """Pluralize a word using the inflect library."""
    engine = inflect.engine()
    return engine.plural(word)

class CustomExtension(Extension):
    """A custom Jinja2 extension to add filters."""
    def __init__(self, environment):
        super().__init__(environment)
        # Register the custom filter
        environment.filters['pluralize'] = pluralize

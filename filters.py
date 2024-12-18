import inflect

def pluralize(word):
    """Pluralize a word using the inflect library."""
    engine = inflect.engine()
    return engine.plural(word)

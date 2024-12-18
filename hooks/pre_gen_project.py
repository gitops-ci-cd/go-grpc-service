import inflect
from jinja2 import Environment, FileSystemLoader

def pluralize(word):
    p = inflect.engine()
    return p.plural(word)

def add_custom_filters():
    from cookiecutter.generate import environment
    env = environment()
    env.filters['pluralize'] = pluralize

if __name__ == "__main__":
    add_custom_filters()

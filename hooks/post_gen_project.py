import inflect
from jinja2 import Environment, FileSystemLoader

def pluralize(word):
    p = inflect.engine()
    return p.plural(word)

def main():
    # Extend Jinja2 environment
    env = Environment(loader=FileSystemLoader('.'))
    env.filters['pluralize'] = pluralize

if __name__ == '__main__':
    main()

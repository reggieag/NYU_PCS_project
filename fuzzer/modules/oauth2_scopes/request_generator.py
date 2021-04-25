import random
import string


class BaseGenerator:
    """
    BaseGenerator is the base class for generators. Generators are used to fill in request bodies, path params, etc etc
    """

    def __init__(self):
        raise NotImplemented

    def integer(self):
        raise NotImplemented

    def number(self):
        raise NotImplemented

    def boolean(self):
        raise NotImplemented

    def string(self):
        raise NotImplemented


class DefaultRandomGenerator(BaseGenerator):
    """
    DefaultRandomGenerator simply generates a random value for reach data type
    """

    def __init__(self):
        pass

    def integer(self, min_val=10, max_val=1000):
        """
        Generates an integer between min_val and max_val
        """
        return random.randint(min_val, max_val)

    def number(self, min_val=10, max_val=1000):
        """
        Generates a float between min_val and max_val
        """
        return random.uniform(min_val, max_val)

    def boolean(self):
        """
        Generates a boolean
        """
        return bool(random.getrandbits(1))

    def string(self, min_val=10, max_val=20):
        """
        Generates a string with length between min_val and max_val
        """
        size = self.integer(min_val, max_val)
        return ''.join(random.choices(
            string.ascii_letters + string.digits, k=size))

#!/usr/bin/env python3
import numpy as np

class Poisson():
    def __init__(self, d=4, x0=4, N=31):
        self.d = d
        self.x0 = x0
        self.N = N
        self._u = np.zeros((self.N*2+1, self.N*2+1))
    def rho(self, x, y):
        d = self.d
        x0 = self.x0
        return np.exp(-((x-x0)**2+y**2)/(d**2)) - np.exp(-((x+x0)**2+y**2)/(d**2))


def ex1():
    pass

ex1()
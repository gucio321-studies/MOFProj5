#!/usr/bin/env python3
import numpy as np

class Poisson():
    def __init__(self, d=4, x0=4, N=31, dx=1):
        self.d = d
        self.x0 = x0
        self.N = N
        self._u = np.zeros((self.N*2+1, self.N*2+1))
        self.dx = 1
    def rho(self, x, y):
        d = self.d
        x0 = self.x0
        return np.exp(-((x-x0)**2+y**2)/(d**2)) - np.exp(-((x+x0)**2+y**2)/(d**2))
    def S(self):
        result = 0
        dx2 = self.dx**2
        for i in range(-self.N+1, self.N): # range(-5, 5) will return [-5 ... 4]
            for j in range(-self.N+1, self.N):
                a1 = .5 * self.u(i,j) * (self.u(i+1, j) + self.u(i-1, j)-2*self.u(i, j))
                a1 /= dx2
                a2 = .5 * self.u(i,j) * (self.u(i, j+1) + self.u(i, j-1) - 2*self.u(i, j))
                a2 /= dx2
                a3 = self.rho(i, j)*self.u(i, j)
                result += (a1 + a2 + a3)*dx2
        return -result


def ex1():
    pass

ex1()
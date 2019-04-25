import time



def is_prime(n):
    if n <= 3:
        return n > 1
    elif n%2==0 or n%3 ==0 :
        return False
    i = 5
    while i*i < n:
        if n%i == 0 or n%(i+2)==0:
            return False
        i += 6
    return True

def generate_primes(n):
    i = 2
    counter = 0
    primes=[]
    while counter < n:
        if is_prime(i):
            primes.append(i)
            counter += 1
        i += 1
    return primes

#icp: is_concatenation_primes
def icp(num1,num2):
    cat1 = int(str(num1)+str(num2))
    cat2 = int(str(num2)+str(num1))
    return is_prime(cat1) and is_prime(cat2)


n = int(input("Number of primes :"))
start = time.time()
PRIMES = generate_primes(n)
end = time.time()
print(end - start)
TWOS=[]
# p1 < p2
for p1 in PRIMES:
    for p2 in PRIMES[PRIMES.index(p1):n]:
        if icp(p1,p2):
            TWOS.append([p1,p2])
end = time.time()
print(end - start)
THREES=[]
for l1 in TWOS:
    for l2 in TWOS:
        if l1[1]==l2[0] and icp(l1[0],l2[1]):
            THREES.append([l1[0],l1[1],l2[1]])
end = time.time()
print(end - start)
FIVES=[]
# [3, 7, 109] and [109,881,3458]
for l1 in THREES:
    for l2 in THREES:
        if l1[2]==l2[0] and icp(l1[1],l2[1]) and icp(l1[0],l2[1]) and icp(l1[0],l2[2]) and icp(l1[1],l2[2]):
            tmp = [l1[0],l1[1],l1[2],l2[1],l2[2]]
            FIVES.append({sum(tmp):tmp})

print(FIVES)
end = time.time()
print(end - start)
from subprocess import Popen, PIPE
import os

sum_e = 0
num_of_threads = 2
data_sizes = ['128', '512', '1024', '4096']
binary = './benchmark'
algo_l = ['HarrisLinkedList', 'HelpOptimalLFList']
percentages = [('50', '50', '0'), ('20', '10', '70'), ('9', '1', '90')]
duration = '5'
for data_size in data_sizes:
    for add_p, rem_p, con_p in percentages:
        for algo in algo_l:
            while num_of_threads <= 64:
                for num in range(0,5):
                    algo_a = '-a='+algo
                    dur_a = '-d='+duration
                    add_a = '-i='+add_p
                    k_a = '-k='+data_size
                    n_a = '-n='+str(num_of_threads)
                    con_a = '-r='+con_p
                    rem_a = '-x='+rem_p
                    p = Popen([binary, algo_a, dur_a, add_a, k_a, n_a, con_a, rem_a], stdin=PIPE, stdout=PIPE, stderr=PIPE)
                    out, err = p.communicate()
                    exitcode = p.returncode
                    sum_e = sum_e + int(int(out.split()[6])/1000000)
                print data_size+' '+add_p+' '+rem_p+' '+con_p+' '+algo+' '+str(num_of_threads)+' '+str(int(sum_e/5))+' '
                num_of_threads *= 2
                sum_e = 0
            num_of_threads = 2

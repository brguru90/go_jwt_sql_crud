import multiprocessing
import uuid,secrets,string,datetime,psycopg2
from psycopg2.pool import ThreadedConnectionPool
from multiprocessing import Pool
import time
 

def random_str(N):
    return ''.join(secrets.choice(string.ascii_uppercase + string.digits) for i in range(N))


count_insert=0

# just give evenly dividable numbers
def sql_populate(DATA_COUNT=10000,BATCH_SIZE=1000):

    def insert_into_tb(i):
        try:
            connection = psycopg2.connect(user="guru",
                                    password="guru",
                                    host="127.0.0.1",
                                    port="5432",
                                    database="jwt4")

            cursor = connection.cursor()
            postgres_insert_query = f"""INSERT INTO users
                                        ("uuid", "name","email", "description", "createdAt","updatedAt") 
                                        VALUES 
                                        ('{str(uuid.uuid4())}','{random_str(10)}','{random_str(5)}@{random_str(5)}.com','{random_str(400)}','{datetime.datetime.now()}','{datetime.datetime.now()}')"""
            cursor.execute(postgres_insert_query)

            connection.commit()
            count = cursor.rowcount
            # print(count, f"{i}th Record inserted successfully into mobile table")
            print("=",end="")
           

        except (Exception, psycopg2.Error) as error:
            print("Failed to connect DB", error)

        finally:
            # closing database connection.
            if  cursor.close():
                cursor.close()
            if connection:
                connection.close()
                # print("PostgreSQL connection is closed")

    def run_in_batch(DATA_COUNT,BATCH_SIZE):
        global count_insert
        while (DATA_COUNT:=DATA_COUNT-BATCH_SIZE)>=0:
            _process=[]
            for i in range(BATCH_SIZE):
                p=multiprocessing.Process(target=insert_into_tb,args=(i,)) 
                p.start()
                _process.append(p)
            for p in _process:
                p.join()
            count_insert+=BATCH_SIZE
            print("\n ====> Remaining",DATA_COUNT)

        print("completed",count_insert)


    # pool = Pool(processes=4)
    # pool.map(insert_into_tb, (range(DATA_COUNT)))
    run_in_batch(DATA_COUNT,BATCH_SIZE)

            

    
if __name__ == '__main__':
    start_time=time.time()
    sql_populate()
    end_time=time.time()
    print("Duration: ",end_time-start_time)
# Braids: Secure pub/sub for ethos

Braids is a project aimed at elevating the security standards of message brokering within Ethos OS by introducing a robust framework that ensures secure, efficient, and reliable data transmission between distributed applications.

### Build
```
make install
```
Note: To run in test/timed mode. Uncomment these two lines in the make file in order to add both the clients to startup
```
#ethosStringEncode /programs/client1    > $(ETHOSROOT)/etc/init/services/client1
#ethosStringEncode /programs/client2    > $(ETHOSROOT)/etc/init/services/client2
```

### Run

Run ethos with the following command
```
sudo -E ethosRun
```

To open the ethos terminal
```
etAl client.ethos
```

After getting access to the ethos terminal. Run the client1 (Pusher)
```
client1
```

Run the client2 (Puller)
```
client2
```

### Logs

To view the detailed logs to know what's being pushed, what's being pulled, the state of the master queue.
Go to `log` directory in the `rootfs`. 

For auth logs: go to `application/braidsAuth`
For borker logs: go to  `application/braidsBorker`
For client logs: go to `test/braidsClient`

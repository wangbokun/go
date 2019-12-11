# linux
git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure --prefix /opt/librdkafka
make
sudo make install


# mac install

brew install librdkafka

go get -v github.com/confluentinc/confluent-kafka-go/kafka
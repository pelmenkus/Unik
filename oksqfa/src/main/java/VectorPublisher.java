package org.example;

// Первое приложение (Запись данных в топик)

import org.eclipse.paho.client.mqttv3.*;

import java.util.Scanner;

public class VectorPublisher {
    private static final String BROKER_URL = "tcp://broker.emqx.io:1883";
    private static final String TOPIC_NAME = "vectors";

    public static void main(String[] args) {
        // Создание экземпляра MQTT клиента
        MqttClient client;
        try {
            client = new MqttClient(BROKER_URL, MqttClient.generateClientId());
            client.connect();

            // Чтение параметров векторов со стандартного потока ввода
            Scanner scanner = new Scanner(System.in);
            System.out.println("Введите возводимую степень");
            String vector1 = scanner.nextLine();
            System.out.println("Введите значение Х");
            String vector2 = scanner.nextLine();

            // Отправка данных в топик
            MqttMessage message1 = new MqttMessage(vector1.getBytes());
            MqttMessage message2 = new MqttMessage(vector2.getBytes());
            client.publish(TOPIC_NAME + "/vector1", message1);
            client.publish(TOPIC_NAME + "/vector2", message2);

            // Закрытие MQTT клиента
            client.disconnect();
            client.close();
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }
}

// Второе приложение (Чтение данных из топика и выполнение вычислений)
package org.example;
// Второе приложение (Чтение данных из топика и выполнение вычислений)

import org.eclipse.paho.client.mqttv3.*;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;

import static java.lang.Math.pow;

public class VectorSubscriber {
    private static final String BROKER_URL = "tcp://broker.emqx.io:1883";
    private static final String TOPIC_NAME = "vectors";

    private static String stepen;
    private static String ex;

    public static void main(String[] args) {
        // Создание экземпляра MQTT клиента
        MqttClient client;
        try {
            client = new MqttClient(BROKER_URL, MqttClient.generateClientId(), new MemoryPersistence());
            client.setCallback(new MqttCallback() {
                @Override
                public void connectionLost(Throwable cause) {
                    System.out.println("Соединение с MQTT брокером потеряно.");
                }

                @Override
                public void messageArrived(String topic, MqttMessage message) {
                    // Обработка полученных сообщений
                    String params = new String(message.getPayload());
                    if (topic.equals(TOPIC_NAME + "/vector1")) {
                        stepen = params;
                        System.out.println("Получена степень: " + params);
                    } else if (topic.equals(TOPIC_NAME + "/vector2")) {
                        ex = params;
                        System.out.println("Получено значение Х: " + params);
                        // Выполнение вычислений на основе векторов
                        System.out.println("Результат вычислений: " + (Integer.valueOf(stepen)*pow(Integer.valueOf(ex),Integer.valueOf(stepen)-1)));
                    }
                }

                @Override
                public void deliveryComplete(IMqttDeliveryToken token) {
                    // Обработка завершения доставки сообщений
                }
            });

            // Подключение к MQTT брокеру и подписка на топик
            client.connect();
            client.subscribe(TOPIC_NAME + "/#", 0);
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }

}

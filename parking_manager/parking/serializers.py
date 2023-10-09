from rest_framework import serializers

class ManagerRegisterSerializer (serializers.Serializer):
    secret_code = serializers.CharField()
    login = serializers.CharField()
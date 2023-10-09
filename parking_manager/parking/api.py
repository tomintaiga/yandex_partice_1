from rest_framework.views import APIView
from rest_framework.response import Response
from parking import serializers
from parking import models
from rest_framework.response import Response
from rest_framework import status
import logging


logger = logging.getLogger(__name__)
secret_code = "AAAAbbbbb"

class RegisterManager(APIView):

    def post(self, request, format=None):
        serializer = serializers.ManagerRegisterSerializer(data=request.data)
        if serializer.is_valid():
            if serializer.data["secret_code"] != secret_code:
                logger.error("Bad secret code")
                return Response(status=status.HTTP_400_BAD_REQUEST)
            
            models.Manager.objects.create(login=serializer.data["login"], parking=request.parking)
            logger.info(f"New manager {serializer.data['login']} created")

            return Response(status=status.HTTP_200_OK)
        
        logger.error(f"Validation errors {serializer.error_messages}")
        return Response(status=status.HTTP_400_BAD_REQUEST)
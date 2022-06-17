from django.db import models
from django.urls import reverse


# Create your models here.
class Contact(models.Model):

    class Meta:
        db_table = "ref_contacts"

    name = models.CharField(max_length=191, null=False)
    email = models.CharField(max_length=191, null=True)
    phone = models.CharField(max_length=191, null=True)
    website = models.CharField(max_length=191, null=True)
    street_address = models.TextField(null=True)
    date_joined = models.DateField(null=True)

    def __str__(self):
        return self.name
from django.db import models

# Create your models here.

class Task(models.Model):
  
  class Meta:
    db_table = "ref_tasks"
  
  title = models.CharField(max_length=200)
  completed = models.BooleanField(default=False, blank=True, null=True)
      
  def __str__(self):
    return self.title
import graphene

from graphene_django import DjangoObjectType

from library.models import Symptom,Disease

class SymptomType(DjangoObjectType):
    class Meta:
        model = Symptom


class DiseaseType(DjangoObjectType):
    class Meta:
        model = Disease


class Query(object):
    all_symptoms = graphene.List(SymptomType)
    all_diseases = graphene.List(DiseaseType)


    def resolve_all_categories(self, info, **kwargs):
        return Category.objects.all()

    def resolve_all_diseases(self, info, **kwargs):
        # We can easily optimize query count in the resolve method
        return Disease.objects.all()
/**
 * Purpose: This File Exists to Hold Variables that would otherwise be hard coded while the backend calls are created.
 */

var COMPANY_GLOBAL_LIST = ["ININ", "Salesforce"];


//Script

var script_dbl =[
  '{{repeat(300,400)}}',
  {
      key: '{{objectId()}}',
      name: '{{firstName()}} {{surname()}}',
      gender: '{{gender()}}',
      gradYear: "{{random('2018', '2019', '2020')}}",
      workStatus: '{{random("EAD", "F1 Visa", "H1 Visa", "J1 Visa", "Permanent Resident", "TN Visa", "US Citizen",  "US Citizen", "US Citizen")}}',
      ethnicity: "{{random('White', 'White','White', 'Hispanic or Latino','Hispanic or Latino', 'Black or African American', 'Native American or American Indian', 'Asian or Pacific Islander')}}",
      grade: '{{floating(0, 5,1) + floating(0, 5, 1 )}}'
  }  
]
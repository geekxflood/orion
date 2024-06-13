# Notes

- Rethink configuration file structure, because the configuration has a `targets` key that define the configuration 
of the system like how to connect and retrieve information, but the endpoint `/targets` is use to expose to prometheus
and I believe that they can be confusion between the two.

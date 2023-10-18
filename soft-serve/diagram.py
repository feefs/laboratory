from diagrams import Diagram, Edge
from diagrams.generic.blank import Blank

from diagrams.aws.general import User
from diagrams.gcp.network import DNS, FirewallRules
from diagrams.gcp.compute import GCE

with Diagram(filename="architecture",
             show=False,
             graph_attr={"splines": "spline"}):
  client = User("Client")
  dns = DNS("Cloud DNS")
  firewall = FirewallRules("Firewall Policies")
  gce = GCE("Compute Engine\n(Running Soft Serve)")

  Blank() - Edge(color="white") - dns
  client << Edge() >> dns
  client - Edge(
      label=
      "ssh <domain>\ngit clone git@<domain>:<repo>.git\nssh <domain> -p <login port>"
  ) >> firewall
  firewall - Edge(label="Allow ports :22\nand <login port>") >> gce

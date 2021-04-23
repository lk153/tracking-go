from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task(3)
    def fetch_product_pb(self):
        self.client.get("/v1/products?limit=10")

    @task(3)
    def fetch_product(self):
        self.client.get("/products?limit=10")

    @task(1)
    def get_home(self):
        self.client.get("/")

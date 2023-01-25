interface Fn {
  id: string;
  name: string;
  image: string;
  environment: {
    [key: string]: string;
  };
  secrets: string[];
  skip_logging: boolean;
  created_time: string;
}

export default Fn;

import 'package:grpc/grpc_web.dart';

import 'src/generated/echo.pb.dart';
import 'src/generated/echo.pbgrpc.dart';

Future<String> request(String endpoint, String payload) async {
  String result;
  final channel = GrpcWebClientChannel.xhr(Uri.parse(endpoint));
  final client = EchoClient(channel);
  final request = GetRequest();
  request.payload = payload;
  try {
    final response = await client.get(request);
    result = response.payload;
  } catch(e) {
    result = e.toString();
  }

  return result;
}
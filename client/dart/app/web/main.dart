import 'dart:convert';
import 'dart:html';
import 'package:http/http.dart' as http;
import 'package:echo/echo.dart';

final output = querySelector('#output');
final InputElement input = querySelector('#payload');
final ButtonElement btn = querySelector('#btn');

void main() async {
  String configData = await http.read('config.json');
  var config = jsonDecode(configData);
  String channelAddress = config['channelAddress'];
  output.text = "Requesting from $channelAddress";
  btn.disabled = true;
  btn.onClick.listen((_) async => output.text = await request(channelAddress, input.value));
  input.onInput.listen((_) => btn.disabled = input.value == "");
}

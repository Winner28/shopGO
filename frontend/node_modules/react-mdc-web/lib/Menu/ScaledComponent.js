'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var transformPropertyName = void 0;

var getTransformPropertyName = function getTransformPropertyName() {
  if (typeof document === 'undefined') {
    return 'transform';
  }
  if (!transformPropertyName) {
    var el = document.createElement('div');
    transformPropertyName = 'transform' in el.style ? 'transform' : '-webkit-transform';
  }
  return transformPropertyName;
};

var propTypes = {
  scaleX: _propTypes2.default.number,
  scaleY: _propTypes2.default.number,
  style: _propTypes2.default.object
};

var ScaledComponent = function ScaledComponent(WrappedComponent) {
  var Transformer = function Transformer(_ref) {
    var scaleX = _ref.scaleX,
        scaleY = _ref.scaleY,
        style = _ref.style,
        otherProps = _objectWithoutProperties(_ref, ['scaleX', 'scaleY', 'style']);

    var transformProperty = getTransformPropertyName();
    return _react2.default.createElement(WrappedComponent, _extends({
      style: _extends(_defineProperty({}, transformProperty, 'scale(' + scaleX + ', ' + scaleY + ')'), style)
    }, otherProps));
  };
  Transformer.propTypes = propTypes;
  return Transformer;
};
exports.default = ScaledComponent;
'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  large: _propTypes2.default.bool
};

var ROOT = 'mdc-card__title';
var LARGE = ROOT + '--large';

var CardTitle = function CardTitle(_ref) {
  var className = _ref.className,
      children = _ref.children,
      large = _ref.large;
  return _react2.default.createElement(
    'h1',
    {
      className: (0, _classnames3.default)(ROOT, _defineProperty({}, LARGE, large), className)
    },
    children
  );
};
CardTitle.propTypes = propTypes;
exports.default = CardTitle;